package userstorage

import (
	"context"

	"github.com/The-System-Guys/login-service/common"
	usermodel "github.com/The-System-Guys/login-service/module/users/model"
	"gorm.io/gorm"
)

func (pg *pgStore) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	db := pg.db.Table(usermodel.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &user, nil
}
