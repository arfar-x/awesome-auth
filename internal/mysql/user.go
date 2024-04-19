package mysql

import "gorm.io/gorm"

type User struct {
	BaseModel
	Name      string `json:"name" gorm:"index; not null"`
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

func (u UserRepo) Get(model BaseModel) bool {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) Create(model BaseModel) bool {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) Update(model BaseModel) bool {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) Delete(model BaseModel) bool {
	//TODO implement me
	panic("implement me")
}
