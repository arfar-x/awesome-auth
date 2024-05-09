package repositories

import (
	"context"
	"time"

	"awesome-auth/internal/domain"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type RepoInterface interface {
	Get(ctx context.Context, model domain.UserDomain) (domain.UserDomain, error)
	Create(ctx context.Context, model domain.UserDomain) (domain.UserDomain, error)
}
