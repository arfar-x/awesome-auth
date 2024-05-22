package repositories

import (
	"context"
	"errors"
	"time"

	"awesome-auth/internal/entities"
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
	FindByUserID(ctx context.Context, model entities.Token) (entities.Token, error)
	Create(ctx context.Context, model entities.Token) (entities.Token, error)
	Delete(ctx context.Context, model entities.Token) (bool, error)
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

func (t TokenRepo) FindByUserID(ctx context.Context, model entities.Token) (entities.Token, error) {
	token := Token{
		UserID: model.UserID,
	}

	result := t.DB.WithContext(ctx).
		Where(token).
		First(&token)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Token{}, gorm.ErrRecordNotFound
		} else {
			panic(err)
		}
	}

	return toTokenEntity(token), nil
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

	return toTokenEntity(*token), nil
}

func (t TokenRepo) Delete(ctx context.Context, model entities.Token) (bool, error) {
	token := Token{
		Value:  model.Value,
		UserID: model.UserID,
	}

	result := t.DB.WithContext(ctx).
		Unscoped().
		Where(token).
		Delete(&t.Token)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, gorm.ErrRecordNotFound
		} else {
			panic(err)
		}
	}

	return true, nil
}

func toTokenEntity(t Token) entities.Token {
	return entities.Token{
		ID:        t.ID,
		UserID:    t.UserID,
		Value:     t.Value,
		ExpiresAt: t.ExpiresAt,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
