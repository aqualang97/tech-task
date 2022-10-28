package services

import (
	"errors"
	"fmt"
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

func (service *DataWebService) InsertNewDataService(parse *models.DataParse) error {
	if parse == nil {
		return errors.New("not found data")
	}
	err := service.store.Data.InsertNewData(parse)
	if err != nil {
		return err
	}
	return nil
}

func (service *DataWebService) GetByTransactionIDService(id int) (*models.DataParse, error) {

	fmt.Println(12312321312321, id)
	if id > 0 {
		data, err := service.store.Data.GetByTransactionID(id)
		if err != nil {
			return nil, err
		}
		fmt.Println(data.TransactionID)
		return data, nil
	}
	return nil, nil
}

func (service *DataWebService) GetByTerminalIDService(id int) (*models.DataParse, error) {
	if id > 0 {
		data, err := service.store.Data.GetByTerminalID(id)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	return nil, nil
}

func (service *DataWebService) GetByPaymentTypeService(paymentType string) (*[]models.DataParse, error) {
	fmt.Println(paymentType, "serv")
	if paymentType != "" {
		data, err := service.store.Data.GetByPaymentType(paymentType)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	return nil, nil
}

func (service *DataWebService) GetByStatusService(status string) (*[]models.DataParse, error) {
	if status != "" {
		data, err := service.store.Data.GetByStatus(status)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	return nil, nil
}

func (service *DataWebService) GetByDataPostService(dateStart, dateEnd int64) (*[]models.DataParse, error) {
	data, err := service.store.Data.GetByDataPost(dateStart, dateEnd)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (service *DataWebService) GetByPaymentNarrativeService(paymentNarrativePart string) (*[]models.DataParse, error) {
	data, err := service.store.Data.GetByPaymentNarrative(paymentNarrativePart)
	if err != nil {
		return nil, err
	}
	return data, nil

}
