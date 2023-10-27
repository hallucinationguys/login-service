package repository

import (
	"github.com/The-System-Guys/login-service.git/pkg/repository/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type pgRepo struct {
	db       *gorm.DB
	userRepo users.UserRepository
}

func NewPGRepo(connectionStr string) UserRepository {
	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &pgRepo{
		db:       db,
		userRepo: users.NewUserRepo(db),
	}
}

func (r pgRepo) AutoMigrate(models ...interface{}) error {
	for idx := range models {
		if err := r.db.AutoMigrate(models[idx]); err != nil {
			return err
		}
	}
	return nil
}

func (r pgRepo) User() users.UserRepository {
	return r.userRepo
}

func (r pgRepo) DB() *gorm.DB {
	return r.db
}
