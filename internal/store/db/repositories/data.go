package repositories

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"sync"
	"techTask/internal/models"
	"techTask/internal/worker_pool"
	"time"
)

type DataRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewDataRepo(db *sql.DB) *DataRepo {
	return &DataRepo{
		DB: db,
	}
}

//func (d *DataRepo) GetByTransactionID(id int) (*models.DataParse, error) {
//	var data models.DataParse
//	var input string
//	var post string
//	err := d.DB.QueryRow(
//		"SELECT transaction_id, request_id, terminal_id, partner_object_id, amount_total, "+
//			"amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, "+
//			"payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bank_mfo, "+
//			"payee_bank_account, payment_narrative FROM data WHERE transaction_id=$1 ORDER BY transaction_id", id).Scan(
//		&data.TransactionID, &data.RequestID, &data.TerminalID, &data.PartnerObjectID,
//		&data.AmountTotal, &data.AmountOriginal, &data.CommissionPS, &data.CommissionClient, &data.CommissionProvider,
//		&input, &post, &data.Status, &data.PaymentType, &data.PaymentNumber, &data.ServiceID,
//		&data.Service, &data.PayeeID, &data.PayeeName, &data.PayeeBankMfo, &data.PayeeBankAccount, &data.PaymentNarrative)
//	i, _ := strconv.ParseInt(input, 10, 64)
//	p, _ := strconv.ParseInt(post, 10, 64)
//
//	data.DateInput = time.Unix(i, 0)
//	data.DatePost = time.Unix(p, 0)
//	if err != nil {
//		return nil, err
//	}
//	return &data, nil
//}
//
//func (d *DataRepo) GetByTerminalID(id int) (*models.DataParse, error) {
//	var data models.DataParse
//	var input string
//	var post string
//	err := d.DB.QueryRow(
//		"SELECT transaction_id, request_id, terminal_id, partner_object_id, amount_total, "+
//			"amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, "+
//			"payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bank_mfo, "+
//			"payee_bank_account, payment_narrative FROM data WHERE terminal_id=$1 ORDER BY transaction_id", id).Scan(
//		&data.TransactionID, &data.RequestID, &data.TerminalID, &data.PartnerObjectID,
//		&data.AmountTotal, &data.AmountOriginal, &data.CommissionPS, &data.CommissionClient, &data.CommissionProvider,
//		&input, &post, &data.Status, &data.PaymentType, &data.PaymentNumber, &data.ServiceID,
//		&data.Service, &data.PayeeID, &data.PayeeName, &data.PayeeBankMfo, &data.PayeeBankAccount, &data.PaymentNarrative)
//
//	i, _ := strconv.ParseInt(input, 10, 64)
//	p, _ := strconv.ParseInt(post, 10, 64)
//
//	data.DateInput = time.Unix(i, 0)
//	data.DatePost = time.Unix(p, 0)
//	if err != nil {
//		fmt.Println(err)
//		return nil, err
//	}
//	return &data, nil
//}
//
//func (d *DataRepo) GetByPaymentType(paymentType string) (*[]models.DataParse, error) {
//
//	var data models.DataParse
//	var listData []models.DataParse
//	var input string
//	var post string
//	rows, err := d.DB.Query(
//		"SELECT transaction_id, request_id, terminal_id, partner_object_id, amount_total, "+
//			"amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, "+
//			"payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bank_mfo, "+
//			"payee_bank_account, payment_narrative FROM data WHERE payment_type=$1 ORDER BY transaction_id", paymentType)
//	if err != nil {
//		return nil, err
//	}
//
//	defer rows.Close()
//	for rows.Next() {
//		err = rows.Scan(
//			&data.TransactionID, &data.RequestID, &data.TerminalID, &data.PartnerObjectID,
//			&data.AmountTotal, &data.AmountOriginal, &data.CommissionPS, &data.CommissionClient, &data.CommissionProvider,
//			&input, &post, &data.Status, &data.PaymentType, &data.PaymentNumber, &data.ServiceID,
//			&data.Service, &data.PayeeID, &data.PayeeName, &data.PayeeBankMfo, &data.PayeeBankAccount, &data.PaymentNarrative)
//		if err != nil {
//			return nil, err
//		}
//		i, _ := strconv.ParseInt(input, 10, 64)
//		p, _ := strconv.ParseInt(post, 10, 64)
//
//		data.DateInput = time.Unix(i, 0)
//		data.DatePost = time.Unix(p, 0)
//		listData = append(listData, data)
//	}
//
//	return &listData, nil
//
//}
//
//func (d *DataRepo) GetByStatus(status string) (*[]models.DataParse, error) {
//	var data models.DataParse
//	var listData []models.DataParse
//	var input string
//	var post string
//	rows, err := d.DB.Query(
//		"SELECT transaction_id, request_id, terminal_id, partner_object_id, amount_total, "+
//			"amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, "+
//			"payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bank_mfo, "+
//			"payee_bank_account, payment_narrative FROM data WHERE status=$1 ORDER BY transaction_id", status)
//	if err != nil {
//		return nil, err
//	}
//
//	defer rows.Close()
//	for rows.Next() {
//		err = rows.Scan(
//			&data.TransactionID, &data.RequestID, &data.TerminalID, &data.PartnerObjectID,
//			&data.AmountTotal, &data.AmountOriginal, &data.CommissionPS, &data.CommissionClient, &data.CommissionProvider,
//			&input, &post, &data.Status, &data.PaymentType, &data.PaymentNumber, &data.ServiceID,
//			&data.Service, &data.PayeeID, &data.PayeeName, &data.PayeeBankMfo, &data.PayeeBankAccount, &data.PaymentNarrative)
//		if err != nil {
//			return nil, err
//		}
//		i, _ := strconv.ParseInt(input, 10, 64)
//		p, _ := strconv.ParseInt(post, 10, 64)
//
//		data.DateInput = time.Unix(i, 0)
//		data.DatePost = time.Unix(p, 0)
//		listData = append(listData, data)
//	}
//
//	return &listData, nil
//}
//
//func (d *DataRepo) GetByDataPost(dateStart, dateEnd int64) (*[]models.DataParse, error) {
//	var data models.DataParse
//	var listData []models.DataParse
//	var input string
//	var post string
//	fmt.Println(dateStart, dateEnd)
//	rows, err := d.DB.Query(
//		"SELECT transaction_id, request_id, terminal_id, partner_object_id, amount_total, "+
//			"amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, "+
//			"payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bank_mfo, "+
//			"payee_bank_account, payment_narrative FROM data WHERE date_post>=$1 AND date_post<=$2 ORDER BY transaction_id", dateStart, dateEnd)
//	if err != nil {
//		return nil, err
//	}
//
//	defer rows.Close()
//	for rows.Next() {
//		err = rows.Scan(
//			&data.TransactionID, &data.RequestID, &data.TerminalID, &data.PartnerObjectID,
//			&data.AmountTotal, &data.AmountOriginal, &data.CommissionPS, &data.CommissionClient, &data.CommissionProvider,
//			&input, &post, &data.Status, &data.PaymentType, &data.PaymentNumber, &data.ServiceID,
//			&data.Service, &data.PayeeID, &data.PayeeName, &data.PayeeBankMfo, &data.PayeeBankAccount, &data.PaymentNarrative)
//		if err != nil {
//			return nil, err
//		}
//
//		i, _ := strconv.ParseInt(input, 10, 64)
//		p, _ := strconv.ParseInt(post, 10, 64)
//
//		data.DateInput = time.Unix(i, 0)
//		data.DatePost = time.Unix(p, 0)
//		listData = append(listData, data)
//	}
//	return &listData, nil
//
//}
//
//func (d *DataRepo) GetByPaymentNarrative(paymentNarrativePart string) (*[]models.DataParse, error) {
//	var data models.DataParse
//	var listData []models.DataParse
//	var input string
//	var post string
//	paymentNarrativePart = "%" + paymentNarrativePart + "%"
//	rows, err := d.DB.Query(
//		"SELECT transaction_id, request_id, terminal_id, partner_object_id, amount_total, "+
//			"amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, "+
//			"payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bank_mfo, "+
//			"payee_bank_account, payment_narrative FROM data WHERE payment_narrative LIKE $1 ORDER BY transaction_id", paymentNarrativePart)
//	fmt.Println(err)
//
//	if err != nil {
//		return nil, err
//	}
//
//	defer rows.Close()
//	for rows.Next() {
//		err = rows.Scan(
//			&data.TransactionID, &data.RequestID, &data.TerminalID, &data.PartnerObjectID,
//			&data.AmountTotal, &data.AmountOriginal, &data.CommissionPS, &data.CommissionClient, &data.CommissionProvider,
//			&input, &post, &data.Status, &data.PaymentType, &data.PaymentNumber, &data.ServiceID,
//			&data.Service, &data.PayeeID, &data.PayeeName, &data.PayeeBankMfo, &data.PayeeBankAccount, &data.PaymentNarrative)
//
//		if err != nil {
//			return nil, err
//		}
//		i, _ := strconv.ParseInt(input, 10, 64)
//		p, _ := strconv.ParseInt(post, 10, 64)
//
//		data.DateInput = time.Unix(i, 0)
//		data.DatePost = time.Unix(p, 0)
//		listData = append(listData, data)
//	}
//
//	return &listData, nil
//}

func (d *DataRepo) Search(queries *models.Queries) (*[]models.DataParse, error) {
	v := reflect.ValueOf(*queries)
	values := make([]interface{}, v.NumField())
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	var dataQ string
	var args []interface{}
	var err error
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
		//fmt.Println(values[i], reflect.TypeOf(values[i]))
	}

	if len(queries.Transactions[0]) > 0 {
		dataQ, args, err = psql.Select("*").From("data").Where(sq.Eq{"transaction_id": values[0]}).ToSql()
		if err != nil {
			return nil, err
		}
	} else if len(queries.Terminal[0]) > 0 {
		dataQ, args, err = psql.Select("*").From("data").Where(sq.Eq{"terminal_id": values[1]}).ToSql()
		if err != nil {
			return nil, err
		}
	} else if len(queries.Status[0]) > 0 {
		dataQ, args, err = psql.Select("*").From("data").Where(sq.Eq{"status": values[2]}).ToSql()
		if err != nil {
			return nil, err
		}
	} else if len(queries.Payment[0]) > 0 {
		dataQ, args, err = psql.Select("*").From("data").Where(sq.Eq{"payment_type": values[3]}).ToSql()
		if err != nil {
			return nil, err
		}
	} else if len(queries.FromDate[0]) > 0 {
		if len(queries.ToDate[0]) > 0 {
			var fromTo []string
			fromTo = append(fromTo, queries.FromDate[0], queries.ToDate[0])

			//var fromTo []string
			//fromTo = append(fromTo, values[4]...)

			dataQ, args, err = psql.Select("*").From("data").Where(sq.Expr("date_post BETWEEN $1 AND $2", queries.FromDate[0], queries.ToDate[0])).ToSql()

			fmt.Println(dataQ)
			if err != nil {
				return nil, err
			}
		}

	} else if len(queries.Narrative[0]) > 0 {
		dataQ, args, err = psql.Select("*").From("data").Where(sq.Eq{"payment_narrative": values[0]}).ToSql()
		if err != nil {
			return nil, err
		}
	}
	fmt.Println(dataQ, args)
	listData, err := d.rowsQuery(dataQ, args)

	return listData, nil
}

func (d *DataRepo) rowsQuery(dataQ string, args []interface{}) (*[]models.DataParse, error) {
	rows, err := d.DB.Query(dataQ, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data models.DataParse
	var listData []models.DataParse
	//var input, post string

	for rows.Next() {
		err = rows.Scan(
			&data.TransactionID, &data.RequestID, &data.TerminalID, &data.PartnerObjectID,
			&data.AmountTotal, &data.AmountOriginal, &data.CommissionPS, &data.CommissionClient, &data.CommissionProvider,
			&data.DateInput, &data.DatePost, &data.Status, &data.PaymentType, &data.PaymentNumber, &data.ServiceID,
			&data.Service, &data.PayeeID, &data.PayeeName, &data.PayeeBankMfo, &data.PayeeBankAccount, &data.PaymentNarrative)

		inputStr := data.DateInput.Format("2006-01-02 15:04:05")
		data.DateInput, err = time.Parse("2006-01-02 15:04:05", inputStr)
		if err != nil {
			return nil, err
		}

		postStr := data.DateInput.Format("2006-01-02 15:04:05")
		data.DateInput, err = time.Parse("2006-01-02 15:04:05", postStr)
		if err != nil {
			return nil, err
		}

		if err != nil {
			return nil, err
		}

		//i, _ := strconv.ParseInt(input, 10, 64)
		//p, _ := strconv.ParseInt(post, 10, 64)
		//
		//data.DateInput = time.Unix(i, 0)
		//data.DatePost = time.Unix(p, 0)
		listData = append(listData, data)
	}
	fmt.Println(listData)
	return &listData, nil

}

func (d *DataRepo) StartParse(filePath string, db *sql.DB) error {
	pool := worker_pool.NewPool(4)
	count := pool.Count
	wg := &sync.WaitGroup{}

	file, err := os.Open(filePath)
	if err != nil {
		log.Println("can't open file; err:", err)
		return err
	}
	readFile, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Println("can't read file; err:", err)
		return err
	}
	for i := 0; i < count; i++ {
		go func(i int) {
			pool.Start(filePath, wg, i, db)
		}(i)
		wg.Add(1)

	}

	for i := 1; i < len(readFile); i++ {
		var data models.DataParse

		data.TransactionID, _ = strconv.Atoi(readFile[i][0])
		data.RequestID, _ = strconv.Atoi(readFile[i][1])
		data.TerminalID, _ = strconv.Atoi(readFile[i][2])
		data.PartnerObjectID, _ = strconv.Atoi(readFile[i][3])
		data.AmountTotal, _ = strconv.Atoi(readFile[i][4])
		data.AmountOriginal, _ = strconv.Atoi(readFile[i][5])
		data.CommissionPS, _ = strconv.ParseFloat(readFile[i][6], 64)
		data.CommissionClient, _ = strconv.Atoi(readFile[i][7])
		data.CommissionProvider, _ = strconv.ParseFloat(readFile[i][8], 64)
		data.DateInput, _ = time.Parse("2006-01-02 15:04:05", readFile[i][9])
		data.DatePost, _ = time.Parse("2006-01-02 15:04:05", readFile[i][10])
		data.Status = readFile[i][11]
		data.PaymentType = readFile[i][12]
		data.PaymentNumber = readFile[i][13]
		data.ServiceID, _ = strconv.Atoi(readFile[i][14])
		data.Service = readFile[i][15]
		data.PayeeID, _ = strconv.Atoi(readFile[i][16])
		data.PayeeName = readFile[i][17]
		data.PayeeBankMfo, _ = strconv.Atoi(readFile[i][18])
		data.PayeeBankAccount = readFile[i][19]
		data.PaymentNarrative = readFile[i][20]

		pool.StartSendData <- &data
	}

	pool.Stop()

	wg.Wait()
	return nil
}

func (d *DataRepo) DownloadFile(url, filename string) error {
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
	fmt.Println(filepath.Abs(filename))
	log.Printf("Data save to file %s\n", filename)
	return nil
}
