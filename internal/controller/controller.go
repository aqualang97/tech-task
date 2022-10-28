package controller

import (
	"techTask/config"
	"techTask/internal/services"
)

type Controller struct {
	Data *DataController
}

func NewController(services *services.Manager, cfg *config.Config) *Controller {
	return &Controller{
		Data: NewDataController(services),
	}
}
