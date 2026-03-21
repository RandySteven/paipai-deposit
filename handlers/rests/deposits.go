package rest_handler

import (
	rest_interfaces "github.com/RandySteven/paipai-deposit/interfaces/handlers/rests"
	usecases_interfaces "github.com/RandySteven/paipai-deposit/interfaces/usecases"
)

type Deposits struct {
	Usecases usecases_interfaces.DepositUsecase
}

func NewDeposits(usecases usecases_interfaces.DepositUsecase) *Deposits {
	return &Deposits{
		Usecases: usecases,
	}
}

var _ rest_interfaces.IDepositRest = &Deposits{}
