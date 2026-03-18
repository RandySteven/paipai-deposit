package repositories

import (
	"context"

	"github.com/RandySteven/go-kopi/entities/models"
	repository_interfaces "github.com/RandySteven/go-kopi/interfaces/repositories"
	mysql_client "github.com/RandySteven/go-kopi/pkg/db"
)

type userRepository struct {
	dbx repository_interfaces.DBX
}

func (u *userRepository) Save(ctx context.Context, entity *models.User) (result *models.User, err error) {
	id, err := mysql_client.Save[models.User](ctx, u.dbx(ctx), ``, entity)
	if err != nil {
		return nil, err
	}
	result = &models.User{
		ID: *id,
	}
	return result, nil
}

func (u *userRepository) FindByID(ctx context.Context, id uint64) (result *models.User, err error) {
	result = &models.User{}
	err = mysql_client.FindByID[models.User](ctx, u.dbx(ctx), ``, id, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userRepository) FindAll(ctx context.Context, skip uint64, take uint64) (result []*models.User, err error) {
	result, err = mysql_client.FindAll[models.User](ctx, u.dbx(ctx), ``)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userRepository) Update(ctx context.Context, entity *models.User) (result *models.User, err error) {
	return
}

func (u *userRepository) DeleteByID(ctx context.Context, id uint64) (err error) {
	return
}

func newUserRepository(dbx repository_interfaces.DBX) *userRepository {
	return &userRepository{
		dbx: dbx,
	}
}

var _ repository_interfaces.UserRepository = &userRepository{}
