package userbiz

import (
	"context"

	"github.com/The-System-Guys/login-service/config"
	usermodel "github.com/The-System-Guys/login-service/internal/module/users/model"
	"github.com/The-System-Guys/login-service/pkg/common"
	"github.com/The-System-Guys/login-service/pkg/components/token"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBusiness struct {
	storeUser  LoginStorage
	tokenMaker token.Maker
	hasher     Hasher
}

func NewLoginBusiness(storeUser LoginStorage, tokenMaker token.Maker, hasher Hasher) *loginBusiness {
	return &loginBusiness{
		storeUser:  storeUser,
		tokenMaker: tokenMaker,
		hasher:     hasher,
	}
}

func (business *loginBusiness) Login(ctx context.Context, data *usermodel.LoginUserRequest) (*usermodel.LoginUserResponse, error) {
	user, err := business.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	config, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	err = business.hasher.VerifyPassword(user.Password, data.Password)

	if err != nil {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	accessToken, accessPayload, err := business.tokenMaker.GenerateToken(user.Email, user.GetRole(), config.AccessTokenDuration)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	userRsp := usermodel.UserResponse{
		LastName:  user.LastName,
		FirstName: user.FirstName,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: *user.CreatedAt,
		UpdateAt:  *user.UpdateAt,
	}

	rsp := usermodel.LoginUserResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
		User:                 userRsp,
	}

	return &rsp, nil
}
