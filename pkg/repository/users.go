package repository

import (
	"github.com/The-System-Guys/login-service.git/pkg/repository/users"
	"gorm.io/gorm"
)

// Repo is the interface that wraps the basic methods for a repository.
type UserRepository interface {
	DB() *gorm.DB
	AutoMigrate(models ...interface{}) error
	User() users.UserRepository
}
