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

// Login godoc
// @Summary Login new user
// @Description login user, returns user and set session
// @Tags Authentication
// @Accept application/json
// @Produce application/json
// @Param user body usermodel.LoginUserRequest true "Login user"
// @Success 200 {object} usermodel.LoginUserResponse
// @Failure 400  {object} usermodel.UserResponse "Error"
// @Router /auth/login [POST]
func Login(appCtx components.AppContext) func(*gin.Context) {
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

		c.JSON(http.StatusOK, account)
	}
}
