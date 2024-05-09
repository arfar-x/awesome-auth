package repositories

import (
	"context"
	"errors"

	"awesome-auth/internal/domain"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Username  string `json:"username" gorm:"index; not null"`
	Email     string `json:"email" gorm:"index; not null"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"-"`
}

type UserRepo struct {
	User User
	DB   *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		User: User{},
		DB:   db,
	}
}

func (u *UserRepo) Get(ctx context.Context, model domain.UserDomain) (domain.UserDomain, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) Create(ctx context.Context, model domain.UserDomain) (domain.UserDomain, error) {
	user := &User{
		Username:  model.Username,
		Email:     model.Email,
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Password:  model.Password,
	}

	result := u.DB.WithContext(ctx).
		Model(user).
		Create(&user)

	if err := result.Error; err != nil {
		panic(err)
	}

	return domain.UserDomain{
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Username:  model.Username,
		Email:     model.Email,
		Password:  model.Password,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (u *UserRepo) Update(ctx context.Context, model any) any {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) Delete(ctx context.Context, model any) any {
	//TODO implement me
	panic("implement me")
}
