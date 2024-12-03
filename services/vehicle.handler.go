package services

import "apolloproject/models"

type Vehicle interface {
	GetVehicleByVIN(*string) (*models.Vehicle, error)
	GetVehicles() ([]*models.Vehicle, error)
	CreateVehicle(*models.Vehicle) error
	UpdateVehicle(*models.Vehicle) error
	DeleteVehicle(*string) error
}
