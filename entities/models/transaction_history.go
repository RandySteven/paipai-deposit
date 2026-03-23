package models

import "time"

type TransactionHistory struct {
	ID                     uint64
	TransactionCode        string
	AccountID              uint64
	BalanceID              uint64
	Amount                 float64
	TransactionType        string
	TransactionAmount      float64
	TransactionDate        time.Time
	TransactionStatus      string
	TransactionDescription string
	TransactionReference   string
	TransactionFee         float64
	TransactionTax         float64
	CreatedAt              time.Time
	UpdatedAt              time.Time
	DeletedAt              *time.Time
}
