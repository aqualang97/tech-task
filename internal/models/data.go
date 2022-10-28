package models

import "time"

type DataParse struct {
	TransactionID   int `json:"transaction_id"`
	RequestID       int `json:"request_id"`
	TerminalID      int `json:"terminal_id"`
	PartnerObjectID int `json:"partner_object_id"`

	AmountTotal        int     `json:"amount_total"`
	AmountOriginal     int     `json:"amount_original"`
	CommissionPS       float64 `json:"commission_ps"`
	CommissionClient   int     `json:"commission_client"`
	CommissionProvider float64 `json:"commission_provider"`

	DateInput time.Time `json:"date_input"`
	DatePost  time.Time `json:"date_post"`

	Status           string `json:"status"`
	PaymentType      string `json:"payment_type"`
	PaymentNumber    string `json:"payment_number"`
	ServiceID        int    `json:"service_id"`
	Service          string `json:"service"`
	PayeeID          int    `json:"payee_id"`
	PayeeName        string `json:"payee_name"`
	PayeeBankMfo     int    `json:"payee_bank_mfo"`
	PayeeBankAccount string `json:"payee_bank_account"`
	PaymentNarrative string `json:"payment_narrative"`
}
