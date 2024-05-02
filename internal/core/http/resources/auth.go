package resources

import "time"

type UserShow struct {
	FirstName string
	LastName  string
	Email     string
	Username  string
	CreatedAt time.Time
}
