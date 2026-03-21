package repository_interfaces

import (
	"context"

	"github.com/RandySteven/paipai-deposit/entities/models"
)

type BalanceRepository interface {
	Saver[models.Balance]
	Finder[models.Balance]
	Updater[models.Balance]
	Deleter[models.Balance]
	FindByAccountID(ctx context.Context, accountID uint64) (result *models.Balance, err error)
}
