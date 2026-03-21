package accounts

import (
	"context"

	"github.com/RandySteven/paipai-deposit/apperror"
)

func (a *accountWorkflow) updateAccountStatus(ctx context.Context, executionData *ExecutionData) (err error) {
	account, err := a.accountRepository.FindByID(ctx, executionData.Account.ID)
	if err != nil {
		return apperror.NewCustomError(apperror.ErrInternalServer, "failed to find account", err)
	}

	account.Status = "ACTIVE"
	account, err = a.accountRepository.Update(ctx, account)
	if err != nil {
		return apperror.NewCustomError(apperror.ErrInternalServer, "failed to save account", err)
	}

	executionData.Account = account
	return nil
}
