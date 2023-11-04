package userstorage

import "gorm.io/gorm"

type pgStore struct {
	db *gorm.DB
}

func NewPGStore(db *gorm.DB) *pgStore {
	return &pgStore{db: db}
}
