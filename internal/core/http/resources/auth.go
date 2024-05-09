package resources

import (
	"time"

	"awesome-auth/internal/domain"
)

type UserShow struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

func UserShowResource(model domain.UserDomain) UserShow {
	return UserShow{
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Email:     model.Email,
		Username:  model.Username,
		CreatedAt: model.CreatedAt,
	}
}
