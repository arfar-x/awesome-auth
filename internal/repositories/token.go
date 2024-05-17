package repositories

import (
	"context"
	"time"

	"awesome-auth/internal/entities"
	"awesome-auth/pkg/jwt"
	"gorm.io/gorm"
)

type Token struct {
	BaseModel
	Value     string
	ExpiresAt time.Time
	UserID    uint
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type TokenRepoInterface interface {
	Get(ctx context.Context, model entities.Token) (entities.Token, error)
	Create(ctx context.Context, model entities.Token) (entities.Token, error)
}

type TokenRepo struct {
	Token Token
	DB    *gorm.DB
}

func NewTokenRepo(db *gorm.DB) *TokenRepo {
	return &TokenRepo{
		Token: Token{},
		DB:    db,
	}
}

func (t TokenRepo) Get(ctx context.Context, model entities.Token) (entities.Token, error) {
	//TODO implement me
	panic("implement me")
}

func (t TokenRepo) Create(ctx context.Context, model entities.Token) (entities.Token, error) {
	token := &Token{
		UserID:    model.UserID,
		Value:     model.Value,
		ExpiresAt: model.ExpiresAt,
	}

	result := t.DB.WithContext(ctx).
		Model(token).
		Create(&token)

	if err := result.Error; err != nil {
		panic(err)
	}

	return toTokenEntity(token), nil
}

func (t TokenRepo) ValidateToken(token string) bool {
	return jwt.Validate(token)
}

func toTokenEntity(t *Token) entities.Token {
	return entities.Token{
		ID:        t.ID,
		UserID:    t.UserID,
		Value:     t.Value,
		ExpiresAt: t.ExpiresAt,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
