package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func addSampleDataToDB() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("vehicledb").Collection("vehicles")

	vehicles := []interface{}{
		bson.M{"vin": "1HGCM82633A123456", "manufacturer_name": "Honda", "description": "Compact Sedan", "horse_power": 150, "model_name": "Civic", "model_year": 2022, "purchase_price": 22000, "fuel_type": "Petrol"},
	}

	_, err = collection.InsertMany(ctx, vehicles)
	if err != nil {
		log.Fatalf("Failed to populate database: %v", err)
	}

	log.Println("Database populated successfully!")
}
