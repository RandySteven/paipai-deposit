package usecases_interfaces

import (
	"context"

	"github.com/RandySteven/paipai-deposit/apperror"
	"github.com/RandySteven/paipai-deposit/entities/payloads/requests"
	"github.com/RandySteven/paipai-deposit/entities/payloads/responses"
)

type DepositUsecase interface {
	CreateAccount(ctx context.Context, request *requests.CreateAccountRequest) (response *responses.CreateAccountResponse, appError *apperror.CustomError)
	GetAccountDetail(ctx context.Context, accountNumber string) (response *responses.AccountDetailResponse, appError *apperror.CustomError)
	GetAccountList(ctx context.Context, cifNumber string) (response *responses.ListAccountsResponse, appError *apperror.CustomError)
	Auth(ctx context.Context, request *requests.AuthRequest) (response *responses.TransferResponse, appError *apperror.CustomError)
	Capture(ctx context.Context, request *requests.CaptureRequest) (response *responses.TransferResponse, appError *apperror.CustomError)
}
