package model

import (
	"time"
)

type Users struct {
	ID        int `gorm:"type:int; primary_key"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"uniqueIndex;not null"`
	GithubID  string    `gorm:"type:varchar(255)"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"type:varchar(255);not null"`
	Photo     string    `gorm:"not null"`
	Verified  bool      `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
