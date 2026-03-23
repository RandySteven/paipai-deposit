package auth

import (
	"github.com/RandySteven/paipai-deposit/apperror"
	"github.com/RandySteven/paipai-deposit/entities/payloads/requests"
	"github.com/RandySteven/paipai-deposit/entities/payloads/responses"
	"go.temporal.io/sdk/workflow"
)

func (a *authWorkflow) auth(ctx workflow.Context, request *requests.AuthRequest) (*responses.TransferResponse, error) {
	executionData := &ExecutionData{
		Request:  request,
		Response: &responses.TransferResponse{},
	}

	err := a.workflowExecution.Execute(ctx, executionData)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to execute workflow", err)
	}

	return executionData.Response, nil
}
