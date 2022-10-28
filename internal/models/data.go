package models

import "time"

type DataParse struct {
	TransactionID   int
	RequestID       int
	TerminalID      int
	PartnerObjectID int

	AmountTotal        int
	AmountOriginal     int
	CommissionPS       float64
	CommissionClient   int
	CommissionProvider float64

	DateInput time.Time
	DatePost  time.Time

	Status           string
	PaymentType      string
	PaymentNumber    string
	ServiceID        int
	Service          string
	PayeeID          int
	PayeeName        string
	PayeeBankMfo     int
	PayeeBankAccount string
	PaymentNarrative string
}
