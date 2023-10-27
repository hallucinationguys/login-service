package routers

import (
	"net/http"
	"time"

	"github.com/The-System-Guys/login-service.git/pkg/controllers"
	"github.com/The-System-Guys/login-service.git/pkg/middleware"
	"github.com/The-System-Guys/login-service.git/pkg/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(userRepository repository.UserRepository, authenticationController *controllers.AuthenticationController, usersController *controllers.UserController) *gin.Engine {
	service := gin.Default()

	service.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	service.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "PUSH", "POST"},
        AllowHeaders:     []string{"Origin"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "Welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	authenticationRouter := router.Group("/authentication")
	authenticationRouter.POST("/register", authenticationController.Register)
	authenticationRouter.POST("/login", authenticationController.Login)

	usersRouter := router.Group("/me")
	usersRouter.GET("", middleware.DeserializeUser(userRepository), usersController.GetUsers)
	

	return service
}
