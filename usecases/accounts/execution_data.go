package accounts

import (
	"github.com/RandySteven/paipai-deposit/entities/models"
	"github.com/RandySteven/paipai-deposit/entities/payloads/responses"
)

type ExecutionData struct {
	CIFNumber string

	Account *models.Account
	Balance *models.Balance

	Response *responses.CreateAccountResponse
	Activity string
}

func (e *ExecutionData) SetActivity(activityName string) {
	e.Activity = activityName
}

func (e *ExecutionData) GetActivity() string {
	return e.Activity
}
