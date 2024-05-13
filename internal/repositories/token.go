package repositories

import (
	"context"
	"time"

	"awesome-auth/internal/domain"
	"awesome-auth/pkg/jwt"
	"gorm.io/gorm"
)

type Token struct {
	BaseModel
	User      User
	UserID    uint
	value     string
	expiresAt time.Time
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

func (t TokenRepo) Get(ctx context.Context, model domain.UserDomain) (domain.UserDomain, error) {
	//TODO implement me
	panic("implement me")
}

func (t TokenRepo) Create(ctx context.Context, model domain.UserDomain) (domain.UserDomain, error) {
	//TODO implement me
	panic("implement me")
}

type Tokenable struct {
	Tokens []Token `json:"-" gorm:"foreignKey:UserID,constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (t Tokenable) CreateToken(ctx context.Context, model domain.UserDomain) string {
	token := jwt.CreateToken(model.Username)

	// TODO: Somehow to create a record containing the token in the database
	//result, err := t.

	return token
}

func (t Tokenable) ValidateToken(ctx context.Context) bool {
	return jwt.Validate("payload")
}
