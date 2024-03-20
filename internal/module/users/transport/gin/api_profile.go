package usergin

import (
	"github.com/gin-gonic/gin"
	"github.com/hallucinationguys/login-service/pkg/common"
	"github.com/hallucinationguys/login-service/pkg/components"
	"net/http"
)

// Profile godoc
// @Summary Profile user
// @Description Profile user, returns user information
// @Tags Authentication
// @Accept application/json
// @Produce application/json
// @Param user body usermodel.LoginUserRequest true "Login user"
// @Success 200 {object} common.successRes
// @Failure 400  {object} common.AppError
// @Router /auth/profile [POST]
func Profile(appCtx components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
