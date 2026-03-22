package accounts

import (
	"context"

	"github.com/RandySteven/paipai-deposit/apperror"
	"github.com/RandySteven/paipai-deposit/entities/payloads/responses"
)

func (a *accountWorkflow) updateAccountStatus(ctx context.Context, executionData *ExecutionData) (*ExecutionData, error) {
	account, err := a.accountRepository.FindByID(ctx, executionData.Account.ID)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to find account", err)
	}

	account.Status = "ACTIVE"
	account, err = a.accountRepository.Update(ctx, account)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to save account", err)
	}

	executionData.Account = account
	executionData.Response = &responses.CreateAccountResponse{
		ID:            executionData.Account.AccountNumber,
		AccountNumber: executionData.Account.AccountNumber,
		CIFNumber:     executionData.Account.CIFNumber,
		Status:        executionData.Account.Status,
		CreatedAt:     executionData.Account.CreatedAt,
		UpdatedAt:     executionData.Account.UpdatedAt,
		DeletedAt:     executionData.Account.DeletedAt,
	}
	return executionData, nil
}
