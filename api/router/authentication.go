package routers

import (
	"github.com/The-System-Guys/login-service/components"
	usergin "github.com/The-System-Guys/login-service/module/users/transport/gin"
	"github.com/gin-gonic/gin"
)

func AuthenticationRoute(router *gin.Engine, appCtx components.AppContext) {
	v1 := router.Group("/v1")
	{
		v1.POST("/auth/login", usergin.Login(appCtx))
		v1.POST("/auth/register", usergin.Register(appCtx))
	}
}
