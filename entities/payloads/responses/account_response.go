package responses

import "time"

type (
	CreateAccountResponse struct {
		ID            string     `json:"id"`
		AccountNumber string     `json:"account_number"`
		CIFNumber     string     `json:"cif_number"`
		Status        string     `json:"status"`
		CreatedAt     time.Time  `json:"created_at"`
		UpdatedAt     time.Time  `json:"updated_at"`
		DeletedAt     *time.Time `json:"deleted_at"`
	}

	AccountListResponse struct {
		AccountNumber string     `json:"account_number"`
		CIFNumber     string     `json:"cif_number"`
		Status        string     `json:"status"`
		CreatedAt     time.Time  `json:"created_at"`
		UpdatedAt     time.Time  `json:"updated_at"`
		DeletedAt     *time.Time `json:"deleted_at"`
	}

	ListAccountsResponse struct {
		CIFNumber string                 `json:"cif_number"`
		Accounts  []*AccountListResponse `json:"accounts"`
	}

	AccountBalanceResponse struct {
		BalanceAmount float64 `json:"balance_amount"`
		Currency      string  `json:"currency"`
	}

	AccountDetailResponse struct {
		AccountNumber string                 `json:"account_number"`
		CIFNumber     string                 `json:"cif_number"`
		Status        string                 `json:"status"`
		Balance       AccountBalanceResponse `json:"balance"`
		CreatedAt     time.Time              `json:"created_at"`
		UpdatedAt     time.Time              `json:"updated_at"`
		DeletedAt     *time.Time             `json:"deleted_at"`
	}
)
