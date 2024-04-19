package mysql

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type RepoInterface interface {
	Get(model BaseModel) bool
	Create(model BaseModel) bool
	Update(model BaseModel) bool
	Delete(model BaseModel) bool
}
