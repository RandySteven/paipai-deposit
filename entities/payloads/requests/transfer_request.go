package requests

type (
	AuthRequest struct {
		IdempotencyKey string  `json:"idempotency_key" mapstructure:"idempotency_key" validate:"required"`
		AccountNumber  string  `json:"account_number" mapstructure:"account_number" validate:"required"`
		Amount         float64 `json:"amount" mapstructure:"amount" validate:"required"`
	}

	CaptureRequest struct {
		IdempotencyKey string  `json:"idempotency_key" mapstructure:"idempotency_key" validate:"required"`
		AccountNumber  string  `json:"account_number" mapstructure:"account_number" validate:"required"`
		Amount         float64 `json:"amount" mapstructure:"amount" validate:"required"`
		TransactionID  string  `json:"transaction_id" mapstructure:"transaction_id" validate:"required"`
	}
)
