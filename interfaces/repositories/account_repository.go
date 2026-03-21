package repository_interfaces

import (
	"context"

	"github.com/RandySteven/paipai-deposit/entities/models"
)

type AccountRepository interface {
	Saver[models.Account]
	Finder[models.Account]
	Updater[models.Account]
	Deleter[models.Account]
	FindByAccountNumber(ctx context.Context, accountNumber string) (result *models.Account, err error)
	FindByCIFNumber(ctx context.Context, cifNumber string) (result []*models.Account, err error)
}
