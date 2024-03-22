package usergin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userbiz "github.com/hallucinationguys/login-service/internal/module/users/biz"
	usermodel "github.com/hallucinationguys/login-service/internal/module/users/model"
	userstorage "github.com/hallucinationguys/login-service/internal/module/users/storage"
	"github.com/hallucinationguys/login-service/pkg/common"
	"github.com/hallucinationguys/login-service/pkg/components"
	"github.com/hallucinationguys/login-service/pkg/components/hasher"
)

// Register godoc
// @Summary Register new user
// @Description Register new user
// @Tags Authentication
// @Accept application/json
// @Produce application/json
// @Param user body usermodel.UserCreate true "Login user"
// @Success 200 {object}  common.successRes
// @Failure 400  {object} common.AppError "Error"
// @Router /auth/register [POST]
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

		userRsp := common.SimpleSuccessResponse(rsp)
		c.JSON(http.StatusOK, userRsp)
	}
}
