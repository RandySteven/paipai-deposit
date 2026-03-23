package auth

import (
	"context"

	"github.com/RandySteven/paipai-deposit/apperror"
)

func (a *authWorkflow) checkAccount(ctx context.Context, executionData *ExecutionData) (*ExecutionData, error) {
	account, err := a.accountRepository.FindByAccountNumber(ctx, executionData.Request.AccountNumber)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to find account", err)
	}

	executionData.Account = account
	executionData.SetActivity(authActivityCheckBalanceInsufficient)
	return executionData, nil
}
