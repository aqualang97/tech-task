package parser

import (
	"techTask/internal/models"
)

type DataParser struct {
	ParseData *[]models.DataParse
}

func NewDataParser() *DataParser {
	return &DataParser{}
}

func (d *DataParser) DownloadFile(url, filename string) error {
	return nil
}
