package models

import "time"

type Account struct {
	ID            uint64
	AccountNumber string
	CIFNumber     string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}
