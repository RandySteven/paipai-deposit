package repository_interfaces

import (
	"context"

	"github.com/RandySteven/paipai-deposit/entities/models"
)

type TransactionHistoryRepository interface {
	Saver[models.TransactionHistory]
	Finder[models.TransactionHistory]
	Updater[models.TransactionHistory]
	Deleter[models.TransactionHistory]
	FindByTransactionCode(ctx context.Context, transactionCode string) (result *models.TransactionHistory, err error)
}
