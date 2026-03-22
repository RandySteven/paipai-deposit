package repositories

import (
	"context"

	"github.com/RandySteven/paipai-deposit/entities/models"
	repository_interfaces "github.com/RandySteven/paipai-deposit/interfaces/repositories"
	mysql_client "github.com/RandySteven/paipai-deposit/pkg/db"
	"github.com/RandySteven/paipai-deposit/queries"
)

type accountRepository struct {
	dbx repository_interfaces.DBX
}

// FindByAccountNumber implements [repository_interfaces.AccountRepository].
func (a *accountRepository) FindByAccountNumber(ctx context.Context, accountNumber string) (result *models.Account, err error) {
	result = new(models.Account)
	err = a.dbx(ctx).QueryRowContext(ctx, queries.SelectAccountByAccountNumber.ToString(), accountNumber).Scan(&result.ID, &result.AccountNumber, &result.CIFNumber, &result.Status, &result.CreatedAt, &result.UpdatedAt, &result.DeletedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindByCIFNumber implements [repository_interfaces.AccountRepository].
func (a *accountRepository) FindByCIFNumber(ctx context.Context, cifNumber string) (result []*models.Account, err error) {
	result = make([]*models.Account, 0)
	rows, err := a.dbx(ctx).QueryContext(ctx, queries.SelectAccountsByCIF.ToString(), cifNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		account := new(models.Account)
		err = rows.Scan(&account.ID, &account.AccountNumber, &account.CIFNumber, &account.Status, &account.CreatedAt, &account.UpdatedAt, &account.DeletedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, account)
	}
	return result, nil
}

// DeleteByID implements [repository_interfaces.AccountRepository].
func (a *accountRepository) DeleteByID(ctx context.Context, id uint64) (err error) {
	_, err = a.dbx(ctx).ExecContext(ctx, queries.DeleteAccountByID.ToString(), id)
	if err != nil {
		return err
	}
	return nil
}

// FindAll implements [repository_interfaces.AccountRepository].
func (a *accountRepository) FindAll(ctx context.Context, skip uint64, take uint64) (result []*models.Account, err error) {
	result = make([]*models.Account, 0)
	rows, err := a.dbx(ctx).QueryContext(ctx, queries.SelectAllAccounts.ToString())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		result = append(result, new(models.Account))
	}
	return result, nil
}

// FindByID implements [repository_interfaces.AccountRepository].
func (a *accountRepository) FindByID(ctx context.Context, id uint64) (result *models.Account, err error) {
	result = new(models.Account)
	err = a.dbx(ctx).QueryRowContext(ctx, string(queries.SelectAccountByID), id).Scan(&result.ID, &result.AccountNumber, &result.CIFNumber, &result.Status, &result.CreatedAt, &result.UpdatedAt, &result.DeletedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Save implements [repository_interfaces.AccountRepository].
func (a *accountRepository) Save(ctx context.Context, entity *models.Account) (result *models.Account, err error) {
	id, err := mysql_client.Save[models.Account](ctx, a.dbx(ctx), queries.InsertAccount, &entity.AccountNumber, &entity.CIFNumber, &entity.Status)
	if err != nil {
		return nil, err
	}
	entity.ID = *id
	return entity, nil
}

// Update implements [repository_interfaces.AccountRepository].
func (a *accountRepository) Update(ctx context.Context, entity *models.Account) (result *models.Account, err error) {
	err = mysql_client.Update[models.Account](ctx, a.dbx(ctx), 
		queries.UpdateAccountByID, entity.AccountNumber, entity.CIFNumber, entity.Status, entity.UpdatedAt, entity.ID)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func NewAccountRepository(dbx repository_interfaces.DBX) repository_interfaces.AccountRepository {
	return &accountRepository{
		dbx: dbx,
	}
}
