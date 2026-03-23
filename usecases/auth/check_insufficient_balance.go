package auth

import (
	"context"

	"github.com/RandySteven/paipai-deposit/apperror"
)

func (a *authWorkflow) checkInsufficientBalance(ctx context.Context, executionData *ExecutionData) (*ExecutionData, error) {
	balance, err := a.balanceRepository.FindByAccountID(ctx, executionData.Account.ID)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to find balance", err)
	}

	if balance.BalanceAmount < executionData.Request.Amount {
		return nil, apperror.NewCustomError(apperror.ErrBadRequest, "insufficient balance", nil)
	}

	executionData.Balance = balance
	executionData.SetActivity(authActivitySaveTransactionHistory)
	return executionData, nil
}
