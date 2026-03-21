package models

import "time"

type Balance struct {
	ID            uint64
	AccountID     uint64
	BalanceAmount float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}
