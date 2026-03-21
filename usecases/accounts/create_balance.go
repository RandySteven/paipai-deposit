package accounts

import (
	"context"

	"github.com/RandySteven/paipai-deposit/apperror"
	"github.com/RandySteven/paipai-deposit/entities/models"
)

func (a *accountWorkflow) createBalance(ctx context.Context, executionData *ExecutionData) (err error) {

	balance := &models.Balance{
		AccountID:     executionData.Account.ID,
		BalanceAmount: 0,
	}

	balance, err = a.balanceRepository.Save(ctx, balance)
	if err != nil {
		return apperror.NewCustomError(apperror.ErrInternalServer, "failed to save balance", err)
	}

	executionData.Balance = balance
	executionData.SetActivity(activityUpdateAccountStatus)
	return nil
}
