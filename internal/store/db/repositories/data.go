package repositories

import (
	"database/sql"
	"techTask/internal/models"
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

	prepare, err := d.DB.Prepare("INSERT INTO data(transaction_id, request_id, terminal_id, partner_object_id, amount_total, amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bank_mfo, payee_bank_account, payment_narrative) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21) ON CONFLICT (transaction_id) DO NOTHING")
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

	return nil

}

func (d *DataRepo) GetByTransactionID(id int) (*DataRepo, error) {

	return nil, nil
}

func (d *DataRepo) GetByTerminalID(id int) (*DataRepo, error) {
	return nil, nil

}

func (d *DataRepo) GetByPaymentType(paymentType string) (*[]DataRepo, error) {
	return nil, nil

}

func (d *DataRepo) GetByStatus() (*[]DataRepo, error) {
	return nil, nil

}

func (d *DataRepo) GetByDataPost() (*[]DataRepo, error) {
	return nil, nil

}

func (d *DataRepo) GetByPaymentNarrative() (*[]DataRepo, error) {
	return nil, nil

}
