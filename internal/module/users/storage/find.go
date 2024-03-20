package userstorage

import (
	"context"

	usermodel "github.com/hallucinationguys/login-service/internal/module/users/model"
	"github.com/hallucinationguys/login-service/pkg/common"
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
