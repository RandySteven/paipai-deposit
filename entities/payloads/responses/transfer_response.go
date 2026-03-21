package responses

import "time"

type (
	TransferResponse struct {
		IdempotencyKey string     `json:"idempotency_key"`
		AccountNumber  string     `json:"account_number"`
		TransactionID  string     `json:"transaction_id"`
		Amount         float64    `json:"amount"`
		Status         string     `json:"status"`
		CreatedAt      time.Time  `json:"created_at"`
		UpdatedAt      time.Time  `json:"updated_at"`
		DeletedAt      *time.Time `json:"deleted_at"`
	}
)
