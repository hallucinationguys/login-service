package usergin

import (
	"net/http"

	userbiz "github.com/The-System-Guys/login-service/internal/module/users/biz"
	usermodel "github.com/The-System-Guys/login-service/internal/module/users/model"
	userstorage "github.com/The-System-Guys/login-service/internal/module/users/storage"
	"github.com/The-System-Guys/login-service/pkg/common"
	"github.com/The-System-Guys/login-service/pkg/components"
	"github.com/The-System-Guys/login-service/pkg/components/hasher"
	"github.com/The-System-Guys/login-service/pkg/components/token"
	"github.com/gin-gonic/gin"
)

func Login(appCtx components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.LoginUserRequest

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenMaker := token.NewJWTMaker(appCtx.SecretKey())

		store := userstorage.NewPGStore(db)
		hash := hasher.NewbcryptHash()

		business := userbiz.NewLoginBusiness(store, tokenMaker, hash, 60*60*24*30)
		account, err := business.Login(c.Request.Context(), &loginUserData)
		if err != nil {
			panic(err)
		}

		rsp := usermodel.LoginUserResponse{
			AccessToken:          account.AccessToken,
			AccessTokenExpiresAt: account.AccessTokenExpiresAt,
			User:                 account.User,
		}

		c.JSON(http.StatusOK, rsp)
	}
}
