package controllers

import (
	"net/http"

	"github.com/The-System-Guys/login-service.git/pkg/data/response"
	"github.com/The-System-Guys/login-service.git/pkg/model"
	"github.com/The-System-Guys/login-service.git/pkg/repository"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository repository.UserRepository
}

func NewUsersController(repository repository.UserRepository) *UserController {
	return &UserController{userRepository: repository}
}

// Get Me		godoc
// @Summary			Get Me
// @Description		Get Me
// @Produce			application/json
// @Tags			Authentication
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success			200 {object} response.Response{}
// @Router			/api/me [GET]
func (controller *UserController) GetUsers(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(*model.Users)

	resp := response.UsersResponse { 
		Username:      currentUser.Name,
		Email:     currentUser.Email,
		Photo:     currentUser.Photo,
		CreatedAt: currentUser.CreatedAt,
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all user data!",
		Data:    resp,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
