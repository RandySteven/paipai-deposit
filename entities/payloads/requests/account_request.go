package requests

type (
	CreateAccountRequest struct {
		CIFNumber      string `json:"cif_number" mapstructure:"cif_number" validate:"required"`
		IdempotencyKey string `json:"idempotency_key" mapstructure:"idempotency_key" validate:"required"`
	}
)
