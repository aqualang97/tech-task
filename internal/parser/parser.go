package parser

import (
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"techTask/internal/models"
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

func (d *DataParser) ParseFromFile(data *models.DataParse, goNum int, db *sql.DB) error {
	err := d.InsertNewData(data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (d *DataParser) InsertNewData(parse *models.DataParse) error {

	prepare, err := d.DB.Prepare("INSERT INTO data(transaction_id, request_id, terminal_id, partner_object_id, " +
		"amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, " +
		"status, payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bank_mfo, " +
		"payee_bank_account, payment_narrative) " +
		"VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21) " +
		"ON CONFLICT (transaction_id) DO NOTHING")
	if err != nil {
		return err
	}
	_, err = prepare.Exec(parse.TransactionID, parse.RequestID, parse.TerminalID, parse.PartnerObjectID,
		parse.AmountTotal, parse.AmountOriginal, parse.CommissionPS, parse.CommissionClient, parse.CommissionProvider,
		parse.DateInput, parse.DatePost, parse.Status, parse.PaymentType, parse.PaymentNumber, parse.ServiceID,
		parse.Service, parse.PayeeID, parse.PayeeName, parse.PayeeBankMfo, parse.PayeeBankAccount, parse.PaymentNarrative)
	if err != nil {
		return err
	}
	// parse.DatePost.Unix(),
	return nil

}
