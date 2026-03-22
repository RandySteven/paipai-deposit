package accounts

import (
	"github.com/RandySteven/paipai-deposit/apperror"
	"github.com/RandySteven/paipai-deposit/entities/payloads/requests"
	"github.com/RandySteven/paipai-deposit/entities/payloads/responses"
	"go.temporal.io/sdk/workflow"
)

func (a *accountWorkflow) createAccount(ctx workflow.Context, request *requests.CreateAccountRequest) (*responses.CreateAccountResponse, error) {
	executionData := &ExecutionData{
		CIFNumber: request.CIFNumber,
		Response:  &responses.CreateAccountResponse{},
	}

	err := a.workflowExecution.Execute(ctx, executionData)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to execute workflow", err)
	}

	if executionData.Account == nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "create account workflow completed without account state", nil)
	}

	return executionData.Response, nil
}
