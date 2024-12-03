package controllers

import (
	"apolloproject/models"
	"apolloproject/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type VehicleController struct {
	VehicleService services.Vehicle
}

func NewVehicleController(vehicleService services.Vehicle) VehicleController {
	return VehicleController{
		VehicleService: vehicleService,
	}
}

func (v *VehicleController) GetVehicleByVIN(ctx *gin.Context) {
	vehicleVIN := ctx.Param("vin")
	vehicle, err := v.VehicleService.GetVehicleByVIN(&vehicleVIN)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, vehicle)
}

func (v *VehicleController) GetVehicles(ctx *gin.Context) {
	vehicles, err := v.VehicleService.GetVehicles()

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, vehicles)
}

func (v *VehicleController) CreateVehicle(ctx *gin.Context) {

	var vehicle models.Vehicle

	if err := ctx.ShouldBindJSON(&vehicle); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := v.VehicleService.CreateVehicle(&vehicle)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "vehicle created successfully"})
}

// PUT /vehicle/:vin
func (v *VehicleController) UpdateVehicle(c *gin.Context) {
	vin := c.Param("vin") // Extract VIN from the URL path
	var updatedVehicle models.Vehicle

	// Change 1: Validate JSON body
	if err := c.BindJSON(&updatedVehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input data"})
		return
	}

	// Set the VIN from the URL parameter
	updatedVehicle.VIN = vin

	// Change 2: Update vehicle using VehicleService
	err := v.VehicleService.UpdateVehicle(&updatedVehicle)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"message": "No vehicle found with the given VIN"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating vehicle"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vehicle updated successfully"}) // Return success response
}

// DELETE /vehicle/:vin
func (vc *VehicleController) DeleteVehicle(c *gin.Context) {
	vin := c.Param("vin") // Extract VIN from the URL path

	// Delete vehicle using VehicleService
	err := vc.VehicleService.DeleteVehicle(&vin)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"message": "No vehicle found with the given VIN"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting vehicle"})
		}
		return
	}

	c.JSON(http.StatusNoContent, nil) // Return 204 No Content as per assignment
}

func (v *VehicleController) RegisterVehicleRoutes(rg *gin.RouterGroup) {
	vehicleRoute := rg.Group("/vehicle")
	vehicleRoute.GET("/:vin", v.GetVehicleByVIN)
	vehicleRoute.GET("/", v.GetVehicles)
	vehicleRoute.POST("/", v.CreateVehicle)
	vehicleRoute.PUT("/:vin", v.UpdateVehicle)
	vehicleRoute.DELETE("/:vin", v.DeleteVehicle)
}
