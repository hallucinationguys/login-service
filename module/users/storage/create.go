package userstorage

import (
	"context"

	"github.com/The-System-Guys/login-service/common"
	usermodel "github.com/The-System-Guys/login-service/module/users/model"
)

func (pg *pgStore) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	db := pg.db.Begin()
	data.PrepareForInsertWithUUID()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
