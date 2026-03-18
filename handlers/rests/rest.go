package rest_handler

import (
	rest_interfaces "github.com/RandySteven/go-kopi/interfaces/handlers/rests"
	"github.com/RandySteven/go-kopi/usecases"
)

type Rests struct {
	UserRest rest_interfaces.IUserRest
}

func NewRests(usecases *usecases.Usecases) *Rests {
	return &Rests{
		UserRest: NewUserRest(usecases.UserUsecase),
	}
}
