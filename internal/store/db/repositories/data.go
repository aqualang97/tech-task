package repositories

import (
	"database/sql"
	"encoding/csv"
	"errors"
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
		dataQ, args, err = psql.Select("*").From("data").Where(sq.Eq{"transaction_id": values[0]}).
			OrderBy("transaction_id").ToSql()
		if err != nil {
			return nil, err
		}
	} else if len(queries.Terminal[0]) > 0 {
		dataQ, args, err = psql.Select("*").From("data").Where(sq.Eq{"terminal_id": values[1]}).
			OrderBy("transaction_id").ToSql()
		if err != nil {
			return nil, err
		}
	} else if len(queries.Status[0]) > 0 {
		dataQ, args, err = psql.Select("*").From("data").Where(sq.Eq{"status": values[2]}).
			OrderBy("transaction_id").ToSql()
		if err != nil {
			return nil, err
		}
	} else if len(queries.Payment[0]) > 0 {
		dataQ, args, err = psql.Select("*").From("data").Where(sq.Eq{"payment_type": values[3]}).
			OrderBy("transaction_id").ToSql()
		if err != nil {
			return nil, err
		}
	} else if len(queries.FromDate[0]) > 0 {
		if len(queries.ToDate[0]) > 0 {
			var fromTo []string
			fromTo = append(fromTo, queries.FromDate[0], queries.ToDate[0])

			//var fromTo []string
			//fromTo = append(fromTo, values[4]...)

			dataQ, args, err = psql.Select("*").From("data").
				Where(sq.Expr("date_post BETWEEN $1 AND $2", queries.FromDate[0], queries.ToDate[0])).
				OrderBy("transaction_id").ToSql()

			fmt.Println(dataQ)
			if err != nil {
				return nil, err
			}
		}

	} else if len(queries.Narrative[0]) > 0 {
		nar := "%" + queries.Narrative[0] + "%"
		fmt.Println(nar)
		dataQ, args, err = psql.Select("*").From("data").Where("payment_narrative LIKE $1", nar).
			OrderBy("transaction_id").ToSql()
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
	fmt.Println(err)

	fmt.Println(args...)
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

func (d *DataRepo) CreateFile(queries *models.Queries) error {
	nameFile := ""
	fmt.Println(queries.Transactions[0] == "", len(queries.Terminal), len(queries.Status), len(queries.Payment), len(queries.FromDate), len(queries.ToDate), queries.Narrative)
	switch {
	case queries.Transactions[0] != "":
		nameFile = "transaction_id.csv"
	case queries.Terminal[0] != "":
		nameFile = "terminal_id.csv"
	case queries.Status[0] != "":
		nameFile = "status.csv"
	case queries.Payment[0] != "":
		nameFile = "payment_type.csv"
	case queries.FromDate[0] != "":
		if queries.ToDate[0] != "" {
			nameFile = "from-to.csv"
		}
	case queries.Narrative[0] != "":
		nameFile = "payment_narrative.csv"
	default:
		return errors.New("invalid data")
	}
	data, err := d.Search(queries)
	if err != nil {
		log.Println(err)
		return err
	}

	firstStr := []string{"TransactionId,RequestId,TerminalId,PartnerObjectId,AmountTotal,AmountOriginal,CommissionPS,CommissionClient,CommissionProvider,DateInput,DatePost,Status,PaymentType,PaymentNumber,ServiceId,Service,PayeeId,PayeeName,PayeeBankMfo,PayeeBankAccount,PaymentNarrative"}

	newFile, err := os.Create(nameFile)
	csvwriter := csv.NewWriter(newFile)
	newData := *data
	for i, _ := range newData {
		inpDate := newData[i].DateInput.Format("2006-01-02 15:04:05")
		postDate := newData[i].DateInput.Format("2006-01-02 15:04:05")

		newStr := []string{
			strconv.Itoa(newData[i].TransactionID),
			strconv.Itoa(newData[i].RequestID),
			strconv.Itoa(newData[i].TerminalID),

			strconv.Itoa(newData[i].PartnerObjectID),
			strconv.Itoa(newData[i].AmountTotal),
			strconv.Itoa(newData[i].AmountOriginal),

			fmt.Sprintf("%2f", newData[i].CommissionPS),
			strconv.Itoa(newData[i].CommissionClient),
			fmt.Sprintf("%2f", newData[i].CommissionProvider),
			inpDate,
			postDate,

			newData[i].Status,
			newData[i].PaymentType,
			newData[i].PaymentNumber,

			strconv.Itoa(newData[i].ServiceID),
			newData[i].Service,
			strconv.Itoa(newData[i].PayeeID),
			newData[i].PayeeName,

			strconv.Itoa(newData[i].PayeeBankMfo),
			newData[i].PayeeBankAccount,
			newData[i].PaymentNarrative,
		}
		if i == 0 {
			_ = csvwriter.Write(firstStr)
		}
		_ = csvwriter.Write(newStr)
	}
	csvwriter.Flush()
	err = newFile.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
