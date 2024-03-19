package usergin

import (
	"github.com/The-System-Guys/login-service/pkg/common"
	"github.com/The-System-Guys/login-service/pkg/components"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Profile godoc
// @Summary Profile user
// @Description Profile user, returns user information
// @Tags Authentication
// @Accept application/json
// @Produce application/json
// @Param user body usermodel.LoginUserRequest true "Login user"
// @Success 200 {object} usermodel.LoginUserResponse
// @Failure 400  {object} usermodel.UserResponse "Error"
// @Router /auth/profile [POST]
func Profile(appCtx components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
