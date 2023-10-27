package services

import "github.com/The-System-Guys/login-service.git/pkg/data/request"

type AuthenticationService interface {
	Login(users request.LoginRequest) (string, error)
	Register(users request.SignUpRequest)
}
