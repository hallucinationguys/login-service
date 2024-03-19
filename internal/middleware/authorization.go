package middleware

import (
	"context"
	"errors"
	"fmt"

	usermodel "github.com/The-System-Guys/login-service/internal/module/users/model"
	"github.com/The-System-Guys/login-service/pkg/common"
	"github.com/The-System-Guys/login-service/pkg/components"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthenStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func RequiredAuth(appCtx components.AppContext, authStore AuthenStore) func(c *gin.Context) {
	tokenMaker := appCtx.GetTokenMaker()

	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrorResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrorResponse(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrorResponse(err))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.ValidationToken(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrorResponse(err))
			return
		}

		user, err := authStore.FindUser(c.Request.Context(), map[string]interface{}{"email": payload.Email})

		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
