package capture

import (
	"context"
	"time"

	"github.com/RandySteven/paipai-deposit/apperror"
	"github.com/RandySteven/paipai-deposit/entities/payloads/requests"
	"github.com/RandySteven/paipai-deposit/entities/payloads/responses"
	repository_interfaces "github.com/RandySteven/paipai-deposit/interfaces/repositories"
	temporal_client "github.com/RandySteven/paipai-deposit/pkg/temporal"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

const (
	captureActivityCheckTransactionHistory = "check_transaction_history_authorized"
	captureActivityBalanceDeduct           = "balance_deduct"
	captureActivityUpdateTransactionStatus = "update_transaction_status"

	sgNoNeed = ""
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
	activityOptions := &workflow.ActivityOptions{
		ScheduleToCloseTimeout: 10 * time.Second,
		StartToCloseTimeout:    10 * time.Second,
		HeartbeatTimeout:       10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    10 * time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    10 * time.Second,
			MaximumAttempts:    3,
		},
	}

	c.workflowExecution.AddTransitionActivityWithOptions(captureActivityCheckTransactionHistory, sgNoNeed, c.checkTransactionStatus, activityOptions, captureActivityBalanceDeduct)
	c.workflowExecution.AddTransitionActivityWithOptions(captureActivityBalanceDeduct, sgNoNeed, c.balanceDeduct, activityOptions, captureActivityUpdateTransactionStatus)
	c.workflowExecution.AddTransitionActivityWithOptions(captureActivityUpdateTransactionStatus, sgNoNeed, c.updateTransactionStatus, activityOptions)

	c.workflowExecution.RegisterWorkflow("Capture", c.capture)
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
