package auth

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
	authActivityCheckAccount             = "check_account"
	authActivityCheckBalanceInsufficient = "check_balance_insufficient"
	authActivitySaveTransactionHistory   = "save_transaction_history"

	sgNoNeed = ""
)

type (
	AuthWorkflow interface {
		Auth(ctx context.Context, request *requests.AuthRequest) (response *responses.TransferResponse, appError *apperror.CustomError)
	}

	authWorkflow struct {
		workflowExecution            temporal_client.WorkflowExecution
		transactionHistoryRepository repository_interfaces.TransactionHistoryRepository
		balanceRepository            repository_interfaces.BalanceRepository
		accountRepository            repository_interfaces.AccountRepository
	}
)

func (a *authWorkflow) registerActivitiesAndWorkflows() {
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

	a.workflowExecution.AddTransitionActivityWithOptions(authActivityCheckAccount, sgNoNeed, a.checkAccount, activityOptions, authActivityCheckBalanceInsufficient)
	a.workflowExecution.AddTransitionActivityWithOptions(authActivityCheckBalanceInsufficient, sgNoNeed, a.checkInsufficientBalance, activityOptions, authActivitySaveTransactionHistory)
	a.workflowExecution.AddTransitionActivityWithOptions(authActivitySaveTransactionHistory, sgNoNeed, a.saveTransactionHistory, activityOptions)

	a.workflowExecution.RegisterWorkflow("Auth", a.auth)
}

func (a *authWorkflow) Auth(ctx context.Context, request *requests.AuthRequest) (response *responses.TransferResponse, appError *apperror.CustomError) {
	return
}

func NewAuthWorkflow(workflowExecution temporal_client.WorkflowExecution,
	transactionHistoryRepository repository_interfaces.TransactionHistoryRepository,
	balanceRepository repository_interfaces.BalanceRepository,
	accountRepository repository_interfaces.AccountRepository,
) AuthWorkflow {
	aw := &authWorkflow{
		workflowExecution:            workflowExecution,
		transactionHistoryRepository: transactionHistoryRepository,
		balanceRepository:            balanceRepository,
		accountRepository:            accountRepository,
	}

	aw.registerActivitiesAndWorkflows()

	return aw
}
