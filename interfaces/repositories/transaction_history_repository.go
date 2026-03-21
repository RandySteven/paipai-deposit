package repository_interfaces

import "github.com/RandySteven/paipai-deposit/entities/models"

type TransactionHistoryRepository interface {
	Saver[models.TransactionHistory]
	Finder[models.TransactionHistory]
	Updater[models.TransactionHistory]
	Deleter[models.TransactionHistory]
}
