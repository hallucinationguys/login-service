package controllers

import (
	"net/http"

	"github.com/The-System-Guys/login-service.git/pkg/data/request"
	"github.com/The-System-Guys/login-service.git/pkg/data/response"
	"github.com/The-System-Guys/login-service.git/pkg/services"
	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	authenticationService services.AuthenticationService
}

func NewAuthenticationController(services services.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{authenticationService: services}
}

// LoginUser		godoc
// @Summary			Login User
// @Description		Login User
// @Param			user body request.LoginRequest true "Login user"
// @Produce			application/json
// @Tags			Authentication
// @Success			200 {object} response.Response{}
// @Router			/api/authentication/login [POST]
func (controller *AuthenticationController) Login(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		panic(err)
	}

	token, err := controller.authenticationService.Login(loginRequest)
	
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid username or password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	resp := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully log in!",
		Data:    resp,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

// RegisterUser		godoc
// @Summary			Register User
// @Description		Register User
// @Param			user body request.SignUpRequest true "Login user"
// @Produce			application/json
// @Tags			Authentication
// @Success			200 {object} response.Response{}
// @Router			/api/authentication/register [POST]
func (controller *AuthenticationController) Register(ctx *gin.Context) {
	createUsersRequest := request.SignUpRequest{}
	err := ctx.ShouldBindJSON(&createUsersRequest)
	if err != nil {
		panic(err)
	}
	controller.authenticationService.Register(createUsersRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created user!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
