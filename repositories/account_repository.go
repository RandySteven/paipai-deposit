package repositories

import (
	"context"
	"github.com/RandySteven/paipai-deposit/entities/models"
	repository_interfaces "github.com/RandySteven/paipai-deposit/interfaces/repositories"
)

type accountRepository struct {
	dbx repository_interfaces.DBX
}

// FindByAccountNumber implements [repository_interfaces.AccountRepository].
func (a *accountRepository) FindByAccountNumber(ctx context.Context, accountNumber string) (result *models.Account, err error) {
	panic("unimplemented")
}

// FindByCIFNumber implements [repository_interfaces.AccountRepository].
func (a *accountRepository) FindByCIFNumber(ctx context.Context, cifNumber string) (result []*models.Account, err error) {
	panic("unimplemented")
}

// DeleteByID implements [repository_interfaces.AccountRepository].
func (a *accountRepository) DeleteByID(ctx context.Context, id uint64) (err error) {
	panic("unimplemented")
}

// FindAll implements [repository_interfaces.AccountRepository].
func (a *accountRepository) FindAll(ctx context.Context, skip uint64, take uint64) (result []*models.Account, err error) {
	panic("unimplemented")
}

// FindByID implements [repository_interfaces.AccountRepository].
func (a *accountRepository) FindByID(ctx context.Context, id uint64) (result *models.Account, err error) {
	panic("unimplemented")
}

// Save implements [repository_interfaces.AccountRepository].
func (a *accountRepository) Save(ctx context.Context, entity *models.Account) (result *models.Account, err error) {
	panic("unimplemented")
}

// Update implements [repository_interfaces.AccountRepository].
func (a *accountRepository) Update(ctx context.Context, entity *models.Account) (result *models.Account, err error) {
	panic("unimplemented")
}

func NewAccountRepository(dbx repository_interfaces.DBX) repository_interfaces.AccountRepository {
	return &accountRepository{
		dbx: dbx,
	}
}
