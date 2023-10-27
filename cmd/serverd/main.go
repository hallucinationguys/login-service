package main

import (
	"log"
	"net/http"

	_ "github.com/The-System-Guys/login-service.git/docs"
	"github.com/The-System-Guys/login-service.git/pkg/config"
	"github.com/The-System-Guys/login-service.git/pkg/controllers"
	"github.com/The-System-Guys/login-service.git/pkg/model"
	"github.com/The-System-Guys/login-service.git/pkg/repository"
	"github.com/The-System-Guys/login-service.git/pkg/routers"
	"github.com/The-System-Guys/login-service.git/pkg/services"
	"github.com/go-playground/validator"
)

// @title 	Login Service API
// @version	1.0
// @description Ecosystem The System Guys API Document

// @host 	localhost:8080
// @BasePath /

func main() {


	init, err := config.LoadConfig()
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	validate := validator.New()

	//Init Repository
	userRepository := repository.NewPGRepo(config.DsnDB(&init))
	userRepository.AutoMigrate(&model.Users{})

	//Init Service
	authenticationService := services.NewAuthenticationService(userRepository, validate)
	authenticationController := controllers.NewAuthenticationController(authenticationService)
	usersController := controllers.NewUsersController(userRepository)

	routes := routers.NewRouter(userRepository, authenticationController, usersController)


	server := &http.Server{
		Addr:    ":" + init.ServerPort,
		Handler: routes,
	}


	err = server.ListenAndServe()
	if err != nil { 
		panic(err)
	}
}
