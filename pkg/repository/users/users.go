package users

import (
	"github.com/The-System-Guys/login-service.git/pkg/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(u model.Users) (*model.Users, error)
	Update(u model.Users) (*model.Users, error)
	FindByID(ID int) (*model.Users, error)
	FindByEmail(email string) (*model.Users, error)
	FindAll() ([]model.Users, error)
	DeleteByID(ID int) error
}

type pgRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepository {
	return &pgRepo{
		DB: db,
	}
}

func (r *pgRepo) Save(u model.Users) (*model.Users, error) {
	return &u, r.DB.Create(&u).Error
}

func (r *pgRepo) Update(u model.Users) (*model.Users, error) {
	return &u, r.DB.Save(&u).Error
}

func (r *pgRepo) FindByID(ID int) (*model.Users, error) {
	u := model.Users{}
	return &u, r.DB.Table("users").Where("id = ?", ID).First(&u).Error
}


func (r *pgRepo) FindByEmail(email string) (*model.Users, error) {
	u := model.Users{}
	return &u, r.DB.Table("users").Where("email = ?", email).First(&u).Error
}

func (r *pgRepo) FindAll() ([]model.Users, error) {
	rs := []model.Users{}
	return rs, r.DB.Table("users").Find(&rs).Error
}

func (r *pgRepo) DeleteByID(ID int) error {
	user := model.Users{}
	return r.DB.Table("users").Where("id = ?", ID).Delete(&user).Error
}
