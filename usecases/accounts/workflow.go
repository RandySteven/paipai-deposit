package accounts

import (
	"context"
	"fmt"

	"github.com/RandySteven/paipai-deposit/apperror"
	"github.com/RandySteven/paipai-deposit/entities/payloads/requests"
	"github.com/RandySteven/paipai-deposit/entities/payloads/responses"
	repository_interfaces "github.com/RandySteven/paipai-deposit/interfaces/repositories"
	temporal_client "github.com/RandySteven/paipai-deposit/pkg/temporal"
)

const (
	activityCheckCifNumber      = "check_cif_number"
	activityCreateAccount       = "create_account"
	activityCreateBalance       = "create_balance"
	activityUpdateAccountStatus = "update_account_status"
)

const (
	sgNoNeed = ""
)

type (
	AccountWorkflow interface {
		CreateAccount(ctx context.Context, request *requests.CreateAccountRequest) (response *responses.CreateAccountResponse, appError *apperror.CustomError)
	}

	accountWorkflow struct {
		workflowExecution temporal_client.WorkflowExecution
		accountRepository repository_interfaces.AccountRepository
		balanceRepository repository_interfaces.BalanceRepository
	}
)

func (a *accountWorkflow) registerActivitiesAndWorkflows() {
	a.workflowExecution.AddTransitionActivityWithOptions(activityCheckCifNumber, sgNoNeed, a.checkCifNumber, nil, activityCreateAccount)
	a.workflowExecution.AddTransitionActivityWithOptions(activityCreateAccount, sgNoNeed, a.initiateAccount, nil, activityCreateBalance)
	a.workflowExecution.AddTransitionActivityWithOptions(activityCreateBalance, sgNoNeed, a.createBalance, nil, activityUpdateAccountStatus)
	a.workflowExecution.AddTransitionActivityWithOptions(activityUpdateAccountStatus, sgNoNeed, a.updateAccountStatus, nil)

	a.workflowExecution.RegisterWorkflow("CreateAccount", a.createAccount)
}

func (a *accountWorkflow) CreateAccount(ctx context.Context, request *requests.CreateAccountRequest) (response *responses.CreateAccountResponse, appError *apperror.CustomError) {

	workflowRun, err := a.workflowExecution.StartWorkflow(ctx, temporal_client.StartWorkflowOptions{
		WorkflowID: fmt.Sprintf("CreateAccount_%s", request.IdempotencyKey),
	}, a.createAccount, request)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to start workflow", err)
	}

	err = a.workflowExecution.GetWorkflowResult(ctx, workflowRun.GetID(), workflowRun.GetRunID(), response)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to get workflow result", err)
	}

	return response, nil
}

func NewAccountWorkflow(workflowExecution temporal_client.WorkflowExecution,
	accountRepository repository_interfaces.AccountRepository,
	balanceRepository repository_interfaces.BalanceRepository,
) AccountWorkflow {
	aw := &accountWorkflow{
		workflowExecution: workflowExecution,
		accountRepository: accountRepository,
		balanceRepository: balanceRepository,
	}

	aw.registerActivitiesAndWorkflows()

	return aw
}
