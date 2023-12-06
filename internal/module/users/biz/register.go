package userbiz

import (
	"context"

	usermodel "github.com/The-System-Guys/login-service/internal/module/users/model"
	"github.com/The-System-Guys/login-service/pkg/common"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	HashPassword(data string) (string, error)
	VerifyPassword(hashedPassword string, candidatePassword string) error
}

type registerBusiness struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBusiness(registerStorage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{
		registerStorage: registerStorage,
		hasher:          hasher,
	}
}
func (business *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := business.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if user != nil {
		return usermodel.ErrEmailExisted
	}

	user, _ = business.registerStorage.FindUser(ctx, map[string]interface{}{"phone": data.Phone})
	if user != nil {
		return usermodel.ErrPhoneExisted
	}

	data.Password, _ = business.hasher.HashPassword(data.Password)

	if err := business.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
