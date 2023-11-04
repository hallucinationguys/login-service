package usergin

import (
	"net/http"

	"github.com/The-System-Guys/login-service/components"
	"github.com/The-System-Guys/login-service/components/hasher"
	userbiz "github.com/The-System-Guys/login-service/module/users/biz"
	usermodel "github.com/The-System-Guys/login-service/module/users/model"
	userstorage "github.com/The-System-Guys/login-service/module/users/storage"
	"github.com/gin-gonic/gin"
)

func Register(appCtx components.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewPGStore(db)
		hash := hasher.NewbcryptHash()
		biz := userbiz.NewRegisterBusiness(store, hash)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		rsp := usermodel.UserResponse{
			LastName:  data.LastName,
			FirstName: data.FirstName,
			Email:     data.Email,
			Role:      data.Role,
			CreatedAt: *data.CreatedAt,
			UpdateAt:  *data.UpdateAt,
		}

		c.JSON(http.StatusOK, rsp)
	}
}
