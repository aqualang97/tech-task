package models

import "time"

type DataParse struct {
	TransactionID   int
	RequestID       int
	TerminalID      int
	PartnerObjectID int

	AmountTotal        float64
	AmountOriginal     float64
	CommissionPS       float64
	CommissionClient   float64
	CommissionProvider float64

	DateInput *time.Time
	DatePost  *time.Time

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
