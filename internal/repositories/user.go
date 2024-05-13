package repositories

import (
	"context"
	"errors"

	"awesome-auth/internal/domain"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Tokenable
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
	result := u.DB.WithContext(ctx).
		Model(u.User).
		Where("username = ?", model.Username).
		First(&u.User)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.UserDomain{}, gorm.ErrRecordNotFound
		} else {
			panic(err)
		}
	}

	return toDomainModel(u.User), nil
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

	return toDomainModel(*user), nil
}

func (u *UserRepo) Update(ctx context.Context, model any) any {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) Delete(ctx context.Context, model any) any {
	//TODO implement me
	panic("implement me")
}

// Turn a repository model object into a domain object.
func toDomainModel(user User) domain.UserDomain {
	return domain.UserDomain{
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
