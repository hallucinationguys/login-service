package services

import (
	"errors"

	"github.com/The-System-Guys/login-service.git/pkg/common"
	"github.com/The-System-Guys/login-service.git/pkg/config"
	"github.com/The-System-Guys/login-service.git/pkg/data/request"
	"github.com/The-System-Guys/login-service.git/pkg/model"
	"github.com/The-System-Guys/login-service.git/pkg/repository"
	"github.com/The-System-Guys/login-service.git/pkg/utils"
	"github.com/go-playground/validator"
)

type AuthenticationServiceImpl struct {
	UsersRepository repository.UserRepository
	Validate        *validator.Validate
}

func NewAuthenticationService(usersRepository repository.UserRepository, validate *validator.Validate) AuthenticationService {
	return &AuthenticationServiceImpl{
		UsersRepository: usersRepository,
		Validate:        validate,
	}
}

// Login implements AuthenticationService
func (a *AuthenticationServiceImpl) Login(users request.LoginRequest) (string, error) {
	new_users, users_err := a.UsersRepository.User().FindByEmail(users.Email)
	if users_err != nil {
		return "", errors.New("invalid username or Password")
	}


	verify_error := common.VerifyPassword(new_users.Password, users.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or Password")
	}
	
	config, _ := config.LoadConfig()
	token, err_token := utils.GenerateToken(config.AccessTokenExpiresIn, new_users.ID, config.AccessTokenPrivateKey)
	if err_token != nil {
		panic(err_token)
	}

	return token, nil
}

func (a *AuthenticationServiceImpl) Register(users request.SignUpRequest) {
	hashedPassword, err := common.HashPassword(users.Password)
	if err != nil {
		panic(err)
	}
	newUser := model.Users{
		Name:     users.Name,
		Email:    users.Email,
		Password: hashedPassword,
	}
	a.UsersRepository.User().Save(newUser)
}
