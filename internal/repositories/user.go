package repositories

import (
	"context"
	"errors"

	"awesome-auth/internal/entities"
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

type UserRepoInterface interface {
	Get(ctx context.Context, model entities.User) (entities.User, error)
	Create(ctx context.Context, model entities.User) (entities.User, error)
	Update(ctx context.Context, model entities.User) (entities.User, error)
	Delete(ctx context.Context, model entities.User) (bool, error)
}

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (u *UserRepo) Get(ctx context.Context, model entities.User) (entities.User, error) {
	user := &User{
		Username: model.Username,
	}

	result := u.DB.WithContext(ctx).
		Where(user).
		First(&user)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.User{}, gorm.ErrRecordNotFound
		} else {
			panic(err)
		}
	}

	return toUserEntity(*user), nil
}

func (u *UserRepo) Create(ctx context.Context, model entities.User) (entities.User, error) {
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

	return toUserEntity(*user), nil
}

func (u *UserRepo) Update(ctx context.Context, model entities.User) (entities.User, error) {
	user := &User{
		Username:  model.Username,
		Email:     model.Email,
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Password:  model.Password,
	}

	result := u.DB.WithContext(ctx).
		Where(user).
		Updates(&user)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.User{}, gorm.ErrRecordNotFound
		} else {
			panic(err)
		}
	}

	return toUserEntity(*user), nil
}

func (u *UserRepo) Delete(ctx context.Context, model entities.User) (bool, error) {
	user := User{
		Username:  model.Username,
		Email:     model.Email,
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Password:  model.Password,
	}

	result := u.DB.WithContext(ctx).
		Where(user).
		Delete(&user)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, gorm.ErrRecordNotFound
		} else {
			panic(err)
		}
	}

	return true, nil
}

// Turn a repository model object into a domain entity.
func toUserEntity(user User) entities.User {
	return entities.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
