package rest_handler

import (
	"net/http"

	rest_interfaces "github.com/RandySteven/go-kopi/interfaces/handlers/rests"
	usecases_interfaces "github.com/RandySteven/go-kopi/interfaces/usecases"
)

type UserRest struct {
	userUsecase usecases_interfaces.UserUsecase
}

func (u *UserRest) RegisterUser(w http.ResponseWriter, r *http.Request) {
}

func (u *UserRest) LoginUser(w http.ResponseWriter, r *http.Request) {
}

func NewUserRest(userUsecase usecases_interfaces.UserUsecase) *UserRest {
	return &UserRest{
		userUsecase: userUsecase,
	}
}

var _ rest_interfaces.IUserRest = &UserRest{}
