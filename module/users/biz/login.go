package userbiz

import (
	"context"
	"time"

	"github.com/The-System-Guys/login-service/common"
	"github.com/The-System-Guys/login-service/components/token"
	usermodel "github.com/The-System-Guys/login-service/module/users/model"
	"github.com/rs/zerolog/log"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBusiness struct {
	storeUser  LoginStorage
	tokenMaker token.Maker
	hasher     Hasher
	expire     int
}

func NewLoginBusiness(storeUser LoginStorage, tokenMaker token.Maker, hasher Hasher, expire int) *loginBusiness {
	return &loginBusiness{
		storeUser:  storeUser,
		tokenMaker: tokenMaker,
		hasher:     hasher,
		expire:     expire,
	}
}

func (business *loginBusiness) Login(ctx context.Context, data *usermodel.LoginUserRequest) (*usermodel.LoginUserResponse, error) {
	user, err := business.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	err = business.hasher.VerifyPassword(user.Password, data.Password)

	if err != nil {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	accessToken, accessPayload, err := business.tokenMaker.GenerateToken(user.Email, user.GetRole(), time.Duration(business.expire))
	if err != nil {
		log.Fatal().Err(err).Msg("Bug here")
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
