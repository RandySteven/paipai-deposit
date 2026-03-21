package repositories

import (
	"context"

	"github.com/RandySteven/paipai-deposit/entities/models"
	repository_interfaces "github.com/RandySteven/paipai-deposit/interfaces/repositories"
)

type balanceRepository struct {
	dbx repository_interfaces.DBX
}

// FindByAccountID implements [repository_interfaces.BalanceRepository].
func (b *balanceRepository) FindByAccountID(ctx context.Context, accountID uint64) (result *models.Balance, err error) {
	panic("unimplemented")
}

func NewBalanceRepository(dbx repository_interfaces.DBX) repository_interfaces.BalanceRepository {
	return &balanceRepository{
		dbx: dbx,
	}
}

// DeleteByID implements [repository_interfaces.BalanceRepository].
func (b *balanceRepository) DeleteByID(ctx context.Context, id uint64) (err error) {
	panic("unimplemented")
}

// FindAll implements [repository_interfaces.BalanceRepository].
func (b *balanceRepository) FindAll(ctx context.Context, skip uint64, take uint64) (result []*models.Balance, err error) {
	panic("unimplemented")
}

// FindByID implements [repository_interfaces.BalanceRepository].
func (b *balanceRepository) FindByID(ctx context.Context, id uint64) (result *models.Balance, err error) {
	panic("unimplemented")
}

// Save implements [repository_interfaces.BalanceRepository].
func (b *balanceRepository) Save(ctx context.Context, entity *models.Balance) (result *models.Balance, err error) {
	panic("unimplemented")
}

// Update implements [repository_interfaces.BalanceRepository].
func (b *balanceRepository) Update(ctx context.Context, entity *models.Balance) (result *models.Balance, err error) {
	panic("unimplemented")
}
