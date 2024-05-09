package domain

import "time"

type UserDomain struct {
	ID        uint
	Username  string
	Email     string
	FirstName string
	LastName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
