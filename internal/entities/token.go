package entities

import "time"

type Token struct {
	ID        uint
	UserID    uint
	Value     string
	ExpiresAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
