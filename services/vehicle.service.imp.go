package services

import (
	"context"
	"errors"

	"apolloproject/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type VehicleHandler struct {
	vehicleCollection *mongo.Collection
	ctx               context.Context
}

// Acts like Constructor to return an interface
func NewVehicleService(ctx context.Context, vehicleCollection *mongo.Collection) *VehicleHandler {
	return &VehicleHandler{
		vehicleCollection: vehicleCollection,
		ctx:               ctx,
	}
}

// Sample VIN
//db.collections.find({vin: "1HGCM82633A123456"})

func (v *VehicleHandler) GetVehicleByVIN(vin *string) (*models.Vehicle, error) {
	var vehicle *models.Vehicle

	query := bson.D{
		{
			Key:   "vin",
			Value: vin,
		},
	}

	err := v.vehicleCollection.FindOne(v.ctx, query).Decode(&vehicle)
	return vehicle, err
}

func (v *VehicleHandler) GetVehicles() ([]*models.Vehicle, error) {
	var vehicles []*models.Vehicle

	cursor, err := v.vehicleCollection.Find(v.ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(v.ctx) {
		var vehicle models.Vehicle
		err := cursor.Decode(&vehicle)

		if err != nil {
			return nil, err
		}

		vehicles = append(vehicles, &vehicle)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(v.ctx)

	if len(vehicles) == 0 {
		return nil, errors.New("no (documents) vehicles found")
	}

	return vehicles, nil
}

func (v *VehicleHandler) CreateVehicle(vehicle *models.Vehicle) error {
	_, err := v.vehicleCollection.InsertOne(v.ctx, vehicle)

	return err
}

func (v *VehicleHandler) UpdateVehicle(vehicle *models.Vehicle) (err error) {
	filter := bson.D{
		primitive.E{
			Key:   "vin",
			Value: vehicle.VIN,
		},
	}

	update := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{
					Key:   "vin",
					Value: vehicle.VIN,
				},
				primitive.E{
					Key:   "manufacturer_name",
					Value: vehicle.Manufacturer,
				},
				primitive.E{
					Key:   "description",
					Value: vehicle.Description,
				},
				primitive.E{
					Key:   "fuel_type",
					Value: vehicle.FuelType,
				},
				primitive.E{
					Key:   "horse_power",
					Value: vehicle.HorsePower,
				},
				primitive.E{
					Key:   "model_name",
					Value: vehicle.ModelName,
				},
				primitive.E{
					Key:   "model_year",
					Value: vehicle.ModelYear,
				},
				primitive.E{
					Key:   "purchase_price",
					Value: vehicle.PurchasePrice,
				},
			},
		},
	}

	result, _ := v.vehicleCollection.UpdateOne(v.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (v *VehicleHandler) DeleteVehicle(vin *string) error {
	filter := bson.D{
		bson.E{
			Key:   "vin",
			Value: vin,
		},
	}

	result, _ := v.vehicleCollection.DeleteOne(v.ctx, filter)

	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}

	return nil
}
