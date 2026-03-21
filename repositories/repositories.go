package repositories

import (
	"context"
	"database/sql"

	repository_interfaces "github.com/RandySteven/paipai-deposit/interfaces/repositories"
)

type Repositories struct {
	AccountRepository repository_interfaces.AccountRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	dbx := func(ctx context.Context) repository_interfaces.Trigger {
		return db
	}

	return &Repositories{
		AccountRepository: NewAccountRepository(dbx),
	}
}
