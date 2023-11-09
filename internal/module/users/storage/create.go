package userstorage

import (
	"context"

	usermodel "github.com/The-System-Guys/login-service/internal/module/users/model"
	"github.com/The-System-Guys/login-service/pkg/common"
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
