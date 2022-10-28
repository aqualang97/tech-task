package services

import "techTask/internal/models"

type DataService interface {
	InsertNewDataService(parse *models.DataParse) error
	GetByTransactionIDService(id int) (*models.DataParse, error)
	GetByTerminalIDService(id int) (*models.DataParse, error)
	GetByPaymentTypeService(paymentType string) (*[]models.DataParse, error)
	GetByStatusService(status string) (*[]models.DataParse, error)
	GetByDataPostService(dateStart, dateEnd int64) (*[]models.DataParse, error)
	GetByPaymentNarrativeService(paymentNarrativePart string) (*[]models.DataParse, error)
}
