package services

import (
	"techTask/internal/models"
	"techTask/internal/store"
)

type DataWebService struct {
	store *store.Store
}

func NewDataWebService(store *store.Store) *DataWebService {
	return &DataWebService{
		store: store,
	}

}

func (service *DataWebService) DownloadFileService(url, filename string) error {
	err := service.store.Data.DownloadFile(url, filename)
	if err != nil {
		return err
	}
	return nil
}

func (service *DataWebService) StartInsertService(path string) error {
	err := service.store.Data.StartParse(path, service.store.Data.DB)
	if err != nil {
		return err
	}
	return nil
}

func (service *DataWebService) SearchService(queries *models.Queries) (*[]models.DataParse, error) {
	data, err := service.store.Data.Search(queries)
	if err != nil {
		return nil, err
	}
	return data, nil

}
