package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vehicle struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	VIN           string             `bson:"vin" json:"vin"`
	Manufacturer  string             `bson:"manufacturer_name" json:"manufacturer"`
	Description   string             `bson:"description,omitempty" json:"description,omitempty"`
	HorsePower    int                `bson:"horse_power" json:"horsePower"`
	ModelName     string             `bson:"model_name" json:"modelName"`
	ModelYear     int                `bson:"model_year" json:"modelYear"`
	PurchasePrice float64            `bson:"purchase_price" json:"purchasePrice"`
	FuelType      string             `bson:"fuel_type" json:"fuelType"`
}
