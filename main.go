package main

import (
	"apolloproject/controllers"
	"apolloproject/services"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server            *gin.Engine
	vehicleService    services.Vehicle
	ctx               context.Context
	vehicleCollection *mongo.Collection
	vehicleController controllers.VehicleController
	mongoClient       *mongo.Client
	err               error
)

func init() {
	ctx = context.TODO()

	mongoConn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoClient, err = mongo.Connect(ctx, mongoConn)
	if err != nil {
		log.Fatal(err)
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection established")

	vehicleCollection = mongoClient.Database("vehicledb").Collection("vehicles")

	_, err = vehicleCollection.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys: bson.M{"vin": 1},
			Options: options.Index().SetUnique(true).SetCollation(
				&options.Collation{Locale: "en", Strength: 2},
			),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	vehicleService = services.NewVehicleService(ctx, vehicleCollection)
	vehicleController = controllers.NewVehicleController(vehicleService)
	server = gin.Default()
}

func main() {
	defer mongoClient.Disconnect(ctx)

	server.GET("/vehicle", vehicleController.GetVehicles)
	server.GET("/vehicle/:vin", vehicleController.GetVehicleByVIN)
	server.POST("/vehicle", vehicleController.CreateVehicle)
	server.PUT("/vehicle/:vin", vehicleController.UpdateVehicle)
	server.DELETE("/vehicle/:vin", vehicleController.DeleteVehicle)

	log.Fatal(server.Run(":8080"))
}
