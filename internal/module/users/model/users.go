package usermodel

import (
	"errors"
	"time"

	"github.com/The-System-Guys/login-service/pkg/common"
	"github.com/google/uuid"
)

const EntityName = "User"

type User struct {
	common.PGModelUUID `json:",inline"`
	Email              string   `json:"email" gorm:"column:email;"`
	Password           string   `json:"password" gorm:"column:password;"`
	LastName           string   `json:"last_name" gorm:"column:last_name;"`
	FirstName          string   `json:"first_name" gorm:"column:first_name;"`
	Gender             int      `json:"gender" gorm:"column:gender;"`
	Phone              string   `json:"phone" gorm:"column:phone;"`
	Image              string   `json:"image" gorm:"column:image;"`
	Role               UserRole `json:"role" gorm:"column:role;"`
}

type UserCreate struct {
	common.PGModelUUID `json:",inline"`
	Email              string   `json:"email" gorm:"column:email;" binding:"required"`
	Password           string   `json:"password" gorm:"column:password;" binding:"required"`
	LastName           string   `json:"last_name" gorm:"column:last_name;" binding:"required"`
	FirstName          string   `json:"first_name" gorm:"column:first_name;" binding:"required"`
	Gender             int      `json:"gender" gorm:"column:gender;" binding:"required"`
	Phone              string   `json:"phone" gorm:"column:phone;" binding:"required"`
	Image              string   `json:"image" gorm:"column:image;"`
	Role               UserRole `json:"role" gorm:"column:role;"`
}

type UserRole int

const (
	RoleUser UserRole = 1 << iota
	RoleAdmin
)

func (role UserRole) String() string {
	switch role {
	case RoleAdmin:
		return "admin"
	default:
		return "user"
	}
}

func (u *User) GetUserId() uuid.UUID {
	return u.ID
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetUserGender() int {
	return u.Gender
}

func (u *User) GetUserPhone() string {
	return u.Phone
}

func (u *User) GetRole() string {
	return u.Role.String()
}

func (User) TableName() string {
	return "users"
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (LoginUserRequest) TableName() string {
	return User{}.TableName()
}

func (u *User) NewUserResponse() UserResponse {
	return UserResponse{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		CreatedAt: *u.CreatedAt,
		UpdateAt:  *u.UpdateAt,
	}
}

type UserResponse struct {
	LastName  string    `json:"last_name"`
	FirstName string    `json:"first_name"`
	Email     string    `json:"email"`
	Role      UserRole  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

type LoginUserRequest struct {
	Email    string `json:"email" form:"email" gorm:"column:email" binding:"required"`
	Password string `json:"password" form:"password" gorm:"column:password" binding:"required"`
}

type LoginUserResponse struct {
	AccessToken          string       `json:"access_token"`
	AccessTokenExpiresAt time.Time    `json:"access_token_expires_at"`
	User                 UserResponse `json:"user"`
}

var (
	ErrEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)

	ErrPhoneExisted = common.NewCustomError(
		errors.New("phone has already existed"),
		"phone has already existed",
		"ErrPhoneExisted",
	)
)
