package capture

import (
	"context"

	"github.com/RandySteven/paipai-deposit/apperror"
	"github.com/RandySteven/paipai-deposit/entities/payloads/requests"
	"github.com/RandySteven/paipai-deposit/entities/payloads/responses"
	repository_interfaces "github.com/RandySteven/paipai-deposit/interfaces/repositories"
	temporal_client "github.com/RandySteven/paipai-deposit/pkg/temporal"
)

type (
	CaptureWorkflow interface {
		Capture(ctx context.Context, request *requests.CaptureRequest) (response *responses.TransferResponse, appError *apperror.CustomError)
	}

	captureWorkflow struct {
		workflowExecution            temporal_client.WorkflowExecution
		transactionHistoryRepository repository_interfaces.TransactionHistoryRepository
		balanceRepository            repository_interfaces.BalanceRepository
		accountRepository            repository_interfaces.AccountRepository
	}
)

func (c *captureWorkflow) registerActivitiesAndWorkflows() {
}

func (c *captureWorkflow) Capture(ctx context.Context, request *requests.CaptureRequest) (response *responses.TransferResponse, appError *apperror.CustomError) {
	return
}

func NewCaptureWorkflow(workflowExecution temporal_client.WorkflowExecution,
	transactionHistoryRepository repository_interfaces.TransactionHistoryRepository,
	balanceRepository repository_interfaces.BalanceRepository,
	accountRepository repository_interfaces.AccountRepository,
) CaptureWorkflow {
	cw := &captureWorkflow{
		workflowExecution:            workflowExecution,
		transactionHistoryRepository: transactionHistoryRepository,
		balanceRepository:            balanceRepository,
		accountRepository:            accountRepository,
	}

	cw.registerActivitiesAndWorkflows()

	return cw
}
