package repositories

import (
	"time"

	"awesome-auth/internal/adapters/auth"
	"gorm.io/gorm"
)

type ModelInterface interface{}

type BaseModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type RepoInterface interface {
	Get(model any) any
	//Create(model context.Context) any
	Create(adapter auth.RegisterAdapter) any
	Update(model any) any
	Delete(model any) any
}
