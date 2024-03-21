package middleware

import (
	"context"
	"errors"
	"fmt"

	usermodel "github.com/hallucinationguys/login-service/internal/module/users/model"
	"github.com/hallucinationguys/login-service/pkg/common"
	"github.com/hallucinationguys/login-service/pkg/components"

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
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.NewUnauthorized(err, "authorization header is not provided", "authorizationHeader"))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.NewUnauthorized(err, "invalid authorization header format", "fields"))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.NewUnauthorized(err, "unsupported authorization type", "authorizationType"))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.ValidationToken(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.NewUnauthorized(err, "token has expired", "ErrTokenExpired"))
			return
		}

		user, err := authStore.FindUser(c.Request.Context(), map[string]interface{}{"email": payload.Email})

		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		userRsp := usermodel.UserResponse{
			LastName:  user.LastName,
			FirstName: user.FirstName,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: *user.CreatedAt,
			UpdateAt:  *user.UpdateAt,
		}

		c.Set(common.CurrentUser, userRsp)
		c.Next()
	}
}
