package auth

import (
	"context"

	"github.com/RandySteven/paipai-deposit/apperror"
	"github.com/RandySteven/paipai-deposit/entities/models"
)

func (a *authWorkflow) saveTransactionHistory(ctx context.Context, executionData *ExecutionData) (*ExecutionData, error) {
	transactionHistory, err := a.transactionHistoryRepository.Save(ctx, &models.TransactionHistory{
		AccountID:       executionData.Account.ID,
		BalanceID:       executionData.Balance.ID,
		Amount:          executionData.Request.Amount,
		TransactionType: "AUTH",
	})
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to save transaction history", err)
	}

	executionData.TransactionHistory = transactionHistory
	return executionData, nil
}
