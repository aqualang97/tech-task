package controller

import (
	"log"
	"path/filepath"
	"techTask/config"
	"techTask/internal/models"
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
func (ctr *DataController) DownloadFile(url, filename string) error {

	err := ctr.services.Data.DownloadFileService(url, filename)

	if err != nil {
		return err
	}
	path, _ := filepath.Abs(filename)
	err = ctr.services.Data.StartInsertService(path)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (ctr *DataController) Search(q *models.Queries) (*[]models.DataParse, error) {
	data, err := ctr.services.Data.SearchService(q)
	return data, err
}
