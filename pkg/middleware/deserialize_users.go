package middleware

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/The-System-Guys/login-service.git/pkg/config"
	"github.com/The-System-Guys/login-service.git/pkg/repository"
	"github.com/The-System-Guys/login-service.git/pkg/utils"
	"github.com/gin-gonic/gin"
)

func DeserializeUser(userRepository repository.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		config, _ := config.LoadConfig()
		sub, err := utils.ValidateToken(token, config.AccessTokenPrivateKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		id, err := strconv.Atoi(fmt.Sprint(sub))
		if err != nil {
			panic(err)
		}

		result, err := userRepository.User().FindByID(id)
		
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
			return
		}

		ctx.Set("currentUser", result)
		ctx.Next()

	}
}