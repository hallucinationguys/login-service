package common

import (
	"time"

	"github.com/google/uuid"
)

type PGModelUUID struct {
	ID        uuid.UUID  `json:"-" gorm:"column:id"`
	Status    int        `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdateAt  *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

type PGModelID struct {
	ID        int        `json:"-" gorm:"column:id"`
	Status    int        `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdateAt  *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (pgModel *PGModelUUID) PrepareForInsertWithUUID() {
	now := time.Now().UTC()
	pgModel.ID = uuid.New()
	pgModel.Status = 1
	pgModel.CreatedAt = &now
	pgModel.UpdateAt = &now
}

func (pgModel *PGModelID) PrepareForInsertWithID() {
	now := time.Now().UTC()
	pgModel.ID = 0
	pgModel.Status = 1
	pgModel.CreatedAt = &now
	pgModel.UpdateAt = &now
}
