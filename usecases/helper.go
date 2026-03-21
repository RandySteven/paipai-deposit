package usecases

import (
	"github.com/RandySteven/paipai-deposit/entities/models"
	"github.com/RandySteven/paipai-deposit/entities/payloads/responses"
)

func (u *usecases) mapAccountToResponse(account *models.Account, balance *models.Balance) *responses.AccountDetailResponse {
	return &responses.AccountDetailResponse{
		AccountNumber: account.AccountNumber,
		CIFNumber:     account.CIFNumber,
		Status:        account.Status,
		Balance: responses.AccountBalanceResponse{
			BalanceAmount: balance.BalanceAmount,
			Currency:      "IDR",
		},
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
		DeletedAt: account.DeletedAt,
	}
}

func (u *usecases) mapAccountListToResponse(cifNumber string, accounts []*models.Account) *responses.ListAccountsResponse {
	out := make([]*responses.AccountListResponse, 0, len(accounts))
	for _, a := range accounts {
		if a == nil {
			continue
		}
		out = append(out, &responses.AccountListResponse{
			AccountNumber: a.AccountNumber,
			CIFNumber:     a.CIFNumber,
			Status:        a.Status,
			CreatedAt:     a.CreatedAt,
			UpdatedAt:     a.UpdatedAt,
			DeletedAt:     a.DeletedAt,
		})
	}
	return &responses.ListAccountsResponse{
		CIFNumber: cifNumber,
		Accounts:  out,
	}
}
