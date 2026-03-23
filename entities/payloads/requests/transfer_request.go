package requests

type (
	AuthRequest struct {
		IdempotencyKey string  `json:"idempotency_key" mapstructure:"idempotency_key" validate:"required"`
		AccountNumber  string  `json:"account_number" mapstructure:"account_number" validate:"required"`
		Amount         float64 `json:"amount" mapstructure:"amount" validate:"required"`
	}

	CaptureRequest struct {
		IdempotencyKey  string  `json:"idempotency_key" mapstructure:"idempotency_key" validate:"required"`
		AccountNumber   string  `json:"account_number" mapstructure:"account_number" validate:"required"`
		Amount          float64 `json:"amount" mapstructure:"amount" validate:"required"`
		TransactionCode string  `json:"transaction_code" mapstructure:"transaction_code" validate:"required"`
	}
)
