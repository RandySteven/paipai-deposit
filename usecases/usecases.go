package usecases

import (
	"context"
	"log"

	"github.com/RandySteven/paipai-deposit/apperror"
	"github.com/RandySteven/paipai-deposit/caches"
	"github.com/RandySteven/paipai-deposit/entities/payloads/requests"
	"github.com/RandySteven/paipai-deposit/entities/payloads/responses"
	repository_interfaces "github.com/RandySteven/paipai-deposit/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/paipai-deposit/interfaces/usecases"
	nsq_client "github.com/RandySteven/paipai-deposit/pkg/nsq"
	temporal_client "github.com/RandySteven/paipai-deposit/pkg/temporal"
	"github.com/RandySteven/paipai-deposit/repositories"
	"github.com/RandySteven/paipai-deposit/usecases/accounts"
	"github.com/RandySteven/paipai-deposit/usecases/auth"
	"github.com/RandySteven/paipai-deposit/usecases/capture"
)

type usecases struct {
	accountWorkflow              accounts.AccountWorkflow
	captureWorkflow              capture.CaptureWorkflow
	authWorkflow                 auth.AuthWorkflow
	accountRepository            repository_interfaces.AccountRepository
	balanceRepository            repository_interfaces.BalanceRepository
	transactionHistoryRepository repository_interfaces.TransactionHistoryRepository
	cache                        caches.Caches
	nsq                          nsq_client.Nsq
	temporal                     temporal_client.Temporal
	workflowExecution            temporal_client.WorkflowExecution
}

// CreateAccount implements [usecases_interfaces.DepositUsecase].
func (u *usecases) CreateAccount(ctx context.Context, request *requests.CreateAccountRequest) (response *responses.CreateAccountResponse, appError *apperror.CustomError) {
	log.Println("CreateAccount", request)
	return u.accountWorkflow.CreateAccount(ctx, request)
}

// GetAccountDetail implements [usecases_interfaces.DepositUsecase].
func (u *usecases) GetAccountDetail(ctx context.Context, accountNumber string) (response *responses.AccountDetailResponse, appError *apperror.CustomError) {
	account, err := u.accountRepository.FindByAccountNumber(ctx, accountNumber)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to find account", err)
	}

	balance, err := u.balanceRepository.FindByAccountID(ctx, account.ID)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to find balance", err)
	}

	result := u.mapAccountToResponse(account, balance)
	return result, nil
}

// GetAccountList implements [usecases_interfaces.DepositUsecase].
func (u *usecases) GetAccountList(ctx context.Context, cifNumber string) (response *responses.ListAccountsResponse, appError *apperror.CustomError) {
	accounts, err := u.accountRepository.FindByCIFNumber(ctx, cifNumber)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to find accounts", err)
	}

	result := u.mapAccountListToResponse(cifNumber, accounts)
	return result, nil
}

func (u *usecases) Auth(ctx context.Context, request *requests.AuthRequest) (response *responses.TransferResponse, appError *apperror.CustomError) {
	// cacheKey := fmt.Sprintf("auth:%s:%s", request.AccountNumber, request.IdempotencyKey)

	// response = &responses.TransferResponse{
	// 	IdempotencyKey: request.IdempotencyKey,
	// 	AccountNumber:  request.AccountNumber,
	// 	Amount:         request.Amount,
	// 	Status:         "AUTH",
	// 	CreatedAt:      time.Now(),
	// 	UpdatedAt:      time.Now(),
	// }

	// err := u.cache.Set(ctx, cacheKey, response)
	// if err != nil {
	// 	return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to set cache", err)
	// }

	// account, err := u.accountRepository.FindByAccountNumber(ctx, request.AccountNumber)
	// if err != nil {
	// 	return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to find account", err)
	// }

	// balance, err := u.balanceRepository.FindByAccountID(ctx, account.ID)
	// if err != nil {
	// 	return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to find balance", err)
	// }

	// transactionHistory, err := u.transactionHistoryRepository.Save(ctx, &models.TransactionHistory{
	// 	AccountID:         account.ID,
	// 	BalanceID:         balance.ID,
	// 	TransactionCode:   request.IdempotencyKey,
	// 	Amount:            request.Amount,
	// 	TransactionType:   "AUTH",
	// 	TransactionAmount: request.Amount,
	// 	TransactionDate:   time.Now(),
	// 	TransactionStatus: "AUTH",
	// })
	// if err != nil {
	// 	return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to save transaction history", err)
	// }

	// err = u.workflowExecution.SignalWorkflow(ctx, "payment-service"+transactionHistory.TransactionCode, "", "auth", response)
	// if err != nil {
	// 	log.Println("failed to signal workflow", err)
	// }

	// response.TransactionCode = transactionHistory.TransactionCode

	return u.authWorkflow.Auth(ctx, request)
}

func (u *usecases) Capture(ctx context.Context, request *requests.CaptureRequest) (response *responses.TransferResponse, appError *apperror.CustomError) {
	// authCacheKey := fmt.Sprintf("auth:%s:%s", request.AccountNumber, request.IdempotencyKey)
	// authCache, err := u.cache.Get(ctx, authCacheKey)
	// if err != nil {
	// 	return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to get cache", err)
	// }

	// if authCache == nil || authCache.(*responses.TransferResponse).Status != "AUTH" {
	// 	return nil, apperror.NewCustomError(apperror.ErrBadRequest, "transaction not found", nil)
	// }

	// transactionHistory, err := u.transactionHistoryRepository.FindByTransactionCode(ctx, request.TransactionCode)
	// if err != nil {
	// 	return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to find transaction history", err)
	// }

	// transactionHistory.TransactionStatus = "CAPTURE"
	// transactionHistory.UpdatedAt = time.Now()
	// _, err = u.transactionHistoryRepository.Update(ctx, transactionHistory)
	// if err != nil {
	// 	return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to update transaction history", err)
	// }

	// balance, err := u.balanceRepository.FindByID(ctx, transactionHistory.BalanceID)
	// if err != nil {
	// 	return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to find balance", err)
	// }

	// balance.BalanceAmount -= transactionHistory.TransactionAmount
	// _, err = u.balanceRepository.Update(ctx, balance)
	// if err != nil {
	// 	return nil, apperror.NewCustomError(apperror.ErrInternalServer, "failed to update balance", err)
	// }

	// err = u.workflowExecution.SignalWorkflow(ctx, "payment-service"+request.TransactionCode, "", "capture", request)
	// if err != nil {
	// 	log.Println("failed to signal workflow", err)
	// }

	return u.captureWorkflow.Capture(ctx, request)
}

func NewUsecases(repositories repositories.Repositories,
	redis caches.Caches,
	nsq nsq_client.Nsq,
	temporal temporal_client.Temporal,
) usecases_interfaces.DepositUsecase {
	workflowExecution := temporal_client.NewWorkflowExecution(temporal)
	us := &usecases{
		accountWorkflow:   accounts.NewAccountWorkflow(workflowExecution, repositories.AccountRepository, repositories.BalanceRepository),
		accountRepository: repositories.AccountRepository,
		balanceRepository: repositories.BalanceRepository,
		cache:             redis,
		nsq:               nsq,
		workflowExecution: workflowExecution,
	}
	return us
}
