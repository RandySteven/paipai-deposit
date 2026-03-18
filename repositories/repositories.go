package repositories

import (
	"context"
	"database/sql"

	repository_interfaces "github.com/RandySteven/go-kopi/interfaces/repositories"
)

type Repositories struct {
	UserRepository repository_interfaces.UserRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	dbx := func(ctx context.Context) repository_interfaces.Trigger {
		return db
	}

	return &Repositories{
		UserRepository: newUserRepository(dbx),
	}
}
