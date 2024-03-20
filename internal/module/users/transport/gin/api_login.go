package usergin

import (
	"github.com/gin-gonic/gin"
	userbiz "github.com/hallucinationguys/login-service/internal/module/users/biz"
	usermodel "github.com/hallucinationguys/login-service/internal/module/users/model"
	userstorage "github.com/hallucinationguys/login-service/internal/module/users/storage"
	"github.com/hallucinationguys/login-service/pkg/common"
	"github.com/hallucinationguys/login-service/pkg/components"
	"github.com/hallucinationguys/login-service/pkg/components/hasher"
	"net/http"
)

// Login godoc
// @Summary Login new user
// @Description login user, returns user and set session
// @Tags Authentication
// @Accept application/json
// @Produce application/json
// @Param user body usermodel.LoginUserRequest true "Login user"
// @Success 200 {object} usermodel.LoginUserResponse
// @Failure 400  {object} common.AppError "Error"
// @Router /auth/login [POST]
func Login(appCtx components.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var loginUserData usermodel.LoginUserRequest

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenMaker := appCtx.GetTokenMaker()

		store := userstorage.NewPGStore(db)
		hash := hasher.NewbcryptHash()

		business := userbiz.NewLoginBusiness(store, tokenMaker, hash)
		account, err := business.Login(c.Request.Context(), &loginUserData)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, account)
	}
}
