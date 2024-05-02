package repositories

import (
	"context"
	"fmt"
	"os"

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

func (u UserRepo) Get(model any) any {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) Create(ctx context.Context) any {

	user := ctx.Value("payload")

	if u, ok := user.(User); ok {
		fmt.Println(u.Email)
	}

	//fmt.Println(u.User)
	//os.Exit(10)
	//u.User.LastName = ctx.Value("LastName")
	//u.User.Email = ctx.Value("Email")
	//u.User.Username = ctx.Value("Username")
	//u.User.Password = password.Make(ctx.Value("Password"))

	res := u.DB.WithContext(ctx).
		Model(u.User).
		Create(&u.User)

	fmt.Println(res)
	os.Exit(1)

	return res
	//panic("implement me")
}

func (u UserRepo) Update(model any) any {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) Delete(model any) any {
	//TODO implement me
	panic("implement me")
}
