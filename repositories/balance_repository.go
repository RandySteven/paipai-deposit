package repositories

import (
	"context"

	"github.com/RandySteven/paipai-deposit/entities/models"
	repository_interfaces "github.com/RandySteven/paipai-deposit/interfaces/repositories"
	mysql_client "github.com/RandySteven/paipai-deposit/pkg/db"
	"github.com/RandySteven/paipai-deposit/queries"
)

type balanceRepository struct {
	dbx repository_interfaces.DBX
}

// FindByAccountID implements [repository_interfaces.BalanceRepository].
func (b *balanceRepository) FindByAccountID(ctx context.Context, accountID uint64) (result *models.Balance, err error) {
	err = mysql_client.FindByID[models.Balance](ctx, b.dbx(ctx), queries.SelectBalancesByAccountID, accountID, result)
	if err != nil {
		return nil, err
	}
	return result, nil
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
	err = mysql_client.FindByID[models.Balance](ctx, b.dbx(ctx), queries.SelectBalanceByID, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Save implements [repository_interfaces.BalanceRepository].
func (b *balanceRepository) Save(ctx context.Context, entity *models.Balance) (result *models.Balance, err error) {
	id, err := mysql_client.Save[models.Balance](ctx, b.dbx(ctx), queries.InsertBalance, entity.AccountID, entity.BalanceAmount)
	if err != nil {
		return nil, err
	}
	entity.ID = *id
	return entity, nil
}

// Update implements [repository_interfaces.BalanceRepository].
func (b *balanceRepository) Update(ctx context.Context, entity *models.Balance) (result *models.Balance, err error) {
	err = mysql_client.Update[models.Balance](ctx, b.dbx(ctx), queries.UpdateBalanceByID, entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}
