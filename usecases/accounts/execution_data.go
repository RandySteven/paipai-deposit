package accounts

import (
	"github.com/RandySteven/paipai-deposit/entities/models"
	"github.com/RandySteven/paipai-deposit/entities/payloads/responses"
	temporal_client "github.com/RandySteven/paipai-deposit/pkg/temporal"
)

type ExecutionData struct {
	CIFNumber string

	Account *models.Account
	Balance *models.Balance

	Response *responses.CreateAccountResponse
	CurrentActivity string
}

var _ temporal_client.NavigatableActivity = &ExecutionData{}

func (e *ExecutionData) SetActivity(activityName string) {
	e.CurrentActivity = activityName
}

func (e *ExecutionData) GetActivity() string {
	return e.CurrentActivity
}
