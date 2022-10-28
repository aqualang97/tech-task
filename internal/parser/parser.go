package parser

import (
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"techTask/internal/models"
	"techTask/internal/store/db/repositories"
)

type DataParser struct {
	ParseData *[]models.DataParse
	DB        *sql.DB
}

func NewDataParser(db *sql.DB) *DataParser {
	return &DataParser{DB: db}
}

func (d *DataParser) DownloadFile(url, filename string) error {
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		log.Println("can't download file with this url: ", url)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("Not OK responce - status:%s, error:%s", resp.Status, err)
		return err
	}
	if resp.Header["Content-Type"][0] != "text/csv" {
		log.Printf("Format file is %s, not %s", resp.Header["Content-Type"][0], "text/csv")
		return err
	}
	read, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Can't read Body", err)
		return err
	}
	err = ioutil.WriteFile(filename, read, 0644)
	if err != nil {
		log.Println("Can't write data to file", err)
		return err
	}
	log.Printf("Data save to file %s\n", filename)
	return nil
}

func ParseFromFile(data *models.DataParse, goNum int, db *sql.DB, TX *sql.Tx) error {
	r := repositories.NewDataRepo(db)
	err := r.InsertNewData(data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (d *DataParser) SaveToDB(filePath string) error {

	return nil
}
