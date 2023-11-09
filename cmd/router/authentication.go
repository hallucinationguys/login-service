package routers

import (
	usergin "github.com/The-System-Guys/login-service/internal/module/users/transport/gin"
	"github.com/The-System-Guys/login-service/pkg/components"
	"github.com/gin-gonic/gin"
)

func AuthenticationRoute(router *gin.Engine, appCtx components.AppContext) {
	v1 := router.Group("/v1")
	{
		v1.POST("/auth/login", usergin.Login(appCtx))
		v1.POST("/auth/register", usergin.Register(appCtx))
	}
}
