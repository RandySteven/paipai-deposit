package usecases

import (
	"context"

	"github.com/RandySteven/go-kopi/apperror"
	"github.com/RandySteven/go-kopi/entities/payloads/requests"
	"github.com/RandySteven/go-kopi/entities/payloads/responses"
	usecases_interfaces "github.com/RandySteven/go-kopi/interfaces/usecases"
)

type userUsecase struct{}

func (u *userUsecase) RegisterUser(ctx context.Context, request *requests.UserRegisterRequest) (response *responses.UserRegisterResponse, appError apperror.CustomError) {
	return
}

func (u *userUsecase) LoginUser(ctx context.Context, request *requests.UserLoginRequest) (response *responses.UserLoginResponse, appError apperror.CustomError) {
	return
}

var _ usecases_interfaces.UserUsecase = &userUsecase{}
