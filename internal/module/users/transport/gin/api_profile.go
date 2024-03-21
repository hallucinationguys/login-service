package usergin

import (
	"github.com/gin-gonic/gin"
	"github.com/hallucinationguys/login-service/pkg/common"
	"github.com/hallucinationguys/login-service/pkg/components"
	"net/http"
)

// Profile godoc
// @Summary      Profile user
// @Description  Get user info
// @Tags         Profile
// @Accept       json
// @Produce      json
// @Success      200  {object}  common.successRes
// @Failure      400  {object}  common.AppError
// @Failure      401  {object}  common.AppError
// @Router /profile [GET]
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
func Profile(appCtx components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
