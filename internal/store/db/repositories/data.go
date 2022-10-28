package repositories

import (
	"database/sql"
	"fmt"
	"strconv"
	"techTask/internal/models"
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

func (d *DataRepo) InsertNewData(parse *models.DataParse) error {

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
		parse.DateInput.Unix(), parse.DatePost.Unix(), parse.Status, parse.PaymentType, parse.PaymentNumber, parse.ServiceID,
		parse.Service, parse.PayeeID, parse.PayeeName, parse.PayeeBankMfo, parse.PayeeBankAccount, parse.PaymentNarrative)
	if err != nil {
		return err
	}

	return nil

}

func (d *DataRepo) GetByTransactionID(id int) (*models.DataParse, error) {
	var data models.DataParse
	var input string
	var post string
	err := d.DB.QueryRow(
		"SELECT transaction_id, request_id, terminal_id, partner_object_id, amount_total, "+
			"amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, "+
			"payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bank_mfo, "+
			"payee_bank_account, payment_narrative FROM data WHERE transaction_id=$1 ORDER BY transaction_id", id).Scan(
		&data.TransactionID, &data.RequestID, &data.TerminalID, &data.PartnerObjectID,
		&data.AmountTotal, &data.AmountOriginal, &data.CommissionPS, &data.CommissionClient, &data.CommissionProvider,
		&input, &post, &data.Status, &data.PaymentType, &data.PaymentNumber, &data.ServiceID,
		&data.Service, &data.PayeeID, &data.PayeeName, &data.PayeeBankMfo, &data.PayeeBankAccount, &data.PaymentNarrative)
	i, _ := strconv.ParseInt(input, 10, 64)
	p, _ := strconv.ParseInt(post, 10, 64)

	data.DateInput = time.Unix(i, 0)
	data.DatePost = time.Unix(p, 0)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (d *DataRepo) GetByTerminalID(id int) (*models.DataParse, error) {
	var data models.DataParse
	var input string
	var post string
	err := d.DB.QueryRow(
		"SELECT transaction_id, request_id, terminal_id, partner_object_id, amount_total, "+
			"amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, "+
			"payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bank_mfo, "+
			"payee_bank_account, payment_narrative FROM data WHERE terminal_id=$1 ORDER BY transaction_id", id).Scan(
		&data.TransactionID, &data.RequestID, &data.TerminalID, &data.PartnerObjectID,
		&data.AmountTotal, &data.AmountOriginal, &data.CommissionPS, &data.CommissionClient, &data.CommissionProvider,
		&input, &post, &data.Status, &data.PaymentType, &data.PaymentNumber, &data.ServiceID,
		&data.Service, &data.PayeeID, &data.PayeeName, &data.PayeeBankMfo, &data.PayeeBankAccount, &data.PaymentNarrative)

	i, _ := strconv.ParseInt(input, 10, 64)
	p, _ := strconv.ParseInt(post, 10, 64)

	data.DateInput = time.Unix(i, 0)
	data.DatePost = time.Unix(p, 0)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &data, nil
}

func (d *DataRepo) GetByPaymentType(paymentType string) (*[]models.DataParse, error) {

	var data models.DataParse
	var listData []models.DataParse
	var input string
	var post string
	rows, err := d.DB.Query(
		"SELECT transaction_id, request_id, terminal_id, partner_object_id, amount_total, "+
			"amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, "+
			"payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bank_mfo, "+
			"payee_bank_account, payment_narrative FROM data WHERE payment_type=$1 ORDER BY transaction_id", paymentType)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&data.TransactionID, &data.RequestID, &data.TerminalID, &data.PartnerObjectID,
			&data.AmountTotal, &data.AmountOriginal, &data.CommissionPS, &data.CommissionClient, &data.CommissionProvider,
			&input, &post, &data.Status, &data.PaymentType, &data.PaymentNumber, &data.ServiceID,
			&data.Service, &data.PayeeID, &data.PayeeName, &data.PayeeBankMfo, &data.PayeeBankAccount, &data.PaymentNarrative)
		if err != nil {
			return nil, err
		}
		i, _ := strconv.ParseInt(input, 10, 64)
		p, _ := strconv.ParseInt(post, 10, 64)

		data.DateInput = time.Unix(i, 0)
		data.DatePost = time.Unix(p, 0)
		listData = append(listData, data)
	}

	return &listData, nil

}

func (d *DataRepo) GetByStatus(status string) (*[]models.DataParse, error) {
	var data models.DataParse
	var listData []models.DataParse
	var input string
	var post string
	rows, err := d.DB.Query(
		"SELECT transaction_id, request_id, terminal_id, partner_object_id, amount_total, "+
			"amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, "+
			"payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bank_mfo, "+
			"payee_bank_account, payment_narrative FROM data WHERE status=$1 ORDER BY transaction_id", status)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&data.TransactionID, &data.RequestID, &data.TerminalID, &data.PartnerObjectID,
			&data.AmountTotal, &data.AmountOriginal, &data.CommissionPS, &data.CommissionClient, &data.CommissionProvider,
			&input, &post, &data.Status, &data.PaymentType, &data.PaymentNumber, &data.ServiceID,
			&data.Service, &data.PayeeID, &data.PayeeName, &data.PayeeBankMfo, &data.PayeeBankAccount, &data.PaymentNarrative)
		if err != nil {
			return nil, err
		}
		i, _ := strconv.ParseInt(input, 10, 64)
		p, _ := strconv.ParseInt(post, 10, 64)

		data.DateInput = time.Unix(i, 0)
		data.DatePost = time.Unix(p, 0)
		listData = append(listData, data)
	}

	return &listData, nil
}

func (d *DataRepo) GetByDataPost(dateStart, dateEnd int64) (*[]models.DataParse, error) {
	var data models.DataParse
	var listData []models.DataParse
	var input string
	var post string
	fmt.Println(dateStart, dateEnd)
	rows, err := d.DB.Query(
		"SELECT transaction_id, request_id, terminal_id, partner_object_id, amount_total, "+
			"amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, "+
			"payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bank_mfo, "+
			"payee_bank_account, payment_narrative FROM data WHERE date_post>=$1 AND date_post<=$2 ORDER BY transaction_id", dateStart, dateEnd)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&data.TransactionID, &data.RequestID, &data.TerminalID, &data.PartnerObjectID,
			&data.AmountTotal, &data.AmountOriginal, &data.CommissionPS, &data.CommissionClient, &data.CommissionProvider,
			&input, &post, &data.Status, &data.PaymentType, &data.PaymentNumber, &data.ServiceID,
			&data.Service, &data.PayeeID, &data.PayeeName, &data.PayeeBankMfo, &data.PayeeBankAccount, &data.PaymentNarrative)
		if err != nil {
			return nil, err
		}

		i, _ := strconv.ParseInt(input, 10, 64)
		p, _ := strconv.ParseInt(post, 10, 64)

		data.DateInput = time.Unix(i, 0)
		data.DatePost = time.Unix(p, 0)
		listData = append(listData, data)
	}
	return &listData, nil

}

func (d *DataRepo) GetByPaymentNarrative(paymentNarrativePart string) (*[]models.DataParse, error) {
	var data models.DataParse
	var listData []models.DataParse
	var input string
	var post string
	paymentNarrativePart = "%" + paymentNarrativePart + "%"
	rows, err := d.DB.Query(
		"SELECT transaction_id, request_id, terminal_id, partner_object_id, amount_total, "+
			"amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, "+
			"payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bank_mfo, "+
			"payee_bank_account, payment_narrative FROM data WHERE payment_narrative LIKE $1 ORDER BY transaction_id", paymentNarrativePart)
	fmt.Println(err)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&data.TransactionID, &data.RequestID, &data.TerminalID, &data.PartnerObjectID,
			&data.AmountTotal, &data.AmountOriginal, &data.CommissionPS, &data.CommissionClient, &data.CommissionProvider,
			&input, &post, &data.Status, &data.PaymentType, &data.PaymentNumber, &data.ServiceID,
			&data.Service, &data.PayeeID, &data.PayeeName, &data.PayeeBankMfo, &data.PayeeBankAccount, &data.PaymentNarrative)

		if err != nil {
			return nil, err
		}
		i, _ := strconv.ParseInt(input, 10, 64)
		p, _ := strconv.ParseInt(post, 10, 64)

		data.DateInput = time.Unix(i, 0)
		data.DatePost = time.Unix(p, 0)
		listData = append(listData, data)
	}

	return &listData, nil
}
