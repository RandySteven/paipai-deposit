package accounts

import (
	"context"

	"github.com/RandySteven/paipai-deposit/apperror"
	"github.com/RandySteven/paipai-deposit/entities/models"
	"github.com/google/uuid"
)

func (a *accountWorkflow) initiateAccount(ctx context.Context, executionData *ExecutionData) (err error) {
	account := &models.Account{
		AccountNumber: uuid.New().String(),
		CIFNumber:     executionData.CIFNumber,
		Status:        "PROCESSING",
	}

	account, err = a.accountRepository.Save(ctx, account)
	if err != nil {
		return apperror.NewCustomError(apperror.ErrInternalServer, "failed to save account", err)
	}

	executionData.Account = account
	executionData.SetActivity(activityCreateBalance)
	return nil
}
