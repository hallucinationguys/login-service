package routers

import (
	"net/http"

	"github.com/hallucinationguys/login-service/internal/middleware"
	userstorage "github.com/hallucinationguys/login-service/internal/module/users/storage"
	usergin "github.com/hallucinationguys/login-service/internal/module/users/transport/gin"
	"github.com/hallucinationguys/login-service/pkg/components"

	"github.com/gin-gonic/gin"
)

func AuthenticationRoute(router *gin.Engine, appCtx components.AppContext) {
	userStore := userstorage.NewPGStore(appCtx.GetMainDBConnection())

	router.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Healthy Check")
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	v1 := router.Group("/v1")
	{
		v1.POST("/auth/login", usergin.Login(appCtx))
		v1.POST("/auth/register", usergin.Register(appCtx))
		v1.GET("/profile", middleware.RequiredAuth(appCtx, userStore), usergin.Profile(appCtx))

	}
}
