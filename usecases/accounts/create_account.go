package accounts

import (
	"github.com/RandySteven/paipai-deposit/apperror"
	"github.com/RandySteven/paipai-deposit/entities/payloads/requests"
	"github.com/RandySteven/paipai-deposit/entities/payloads/responses"
	"go.temporal.io/sdk/workflow"
)

func (a *accountWorkflow) createAccount(ctx workflow.Context, request *requests.CreateAccountRequest) (response *responses.CreateAccountResponse, appError *apperror.CustomError) {
	executionData := &ExecutionData{
		CIFNumber: request.CIFNumber,
	}

	err := a.workflowExecution.Execute(ctx, executionData)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to execute workflow", err)
	}

	executionData.Response = &responses.CreateAccountResponse{
		ID: executionData.Account.AccountNumber,
	}

	return executionData.Response, nil
}
