package capture

import (
	"context"

	"github.com/RandySteven/paipai-deposit/apperror"
)

func (c *captureWorkflow) balanceDeduct(ctx context.Context, executionData *ExecutionData) (*ExecutionData, error) {
	balance, err := c.balanceRepository.FindByID(ctx, executionData.TransactionHistory.BalanceID)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to find balance", err)
	}

	balance.BalanceAmount -= executionData.TransactionHistory.TransactionAmount
	_, err = c.balanceRepository.Update(ctx, balance)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to update balance", err)
	}
	return executionData, nil
}
