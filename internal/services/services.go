package services

import "techTask/internal/models"

type DataService interface {
	DownloadFileService(url, filename string) error
	SearchService(queries *models.Queries) (*[]models.DataParse, error)
	StartInsertService(path string) error
	CreateFileService(queries *models.Queries) error
}
