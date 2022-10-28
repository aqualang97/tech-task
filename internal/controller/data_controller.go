package controller

import (
	"techTask/config"
	"techTask/internal/services"
)

type DataController struct {
	services *services.Manager
	cfg      *config.Config
}

func NewDataController(services *services.Manager) *DataController {
	return &DataController{
		services: services,
	}
}
