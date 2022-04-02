package domain

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	Id        uuid.UUID
	Name      string
	Users     []*User
	Comments  []*Comment
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Comment struct {
	Id        uuid.UUID
	Comment   string
	User      *User
	CreatedAt time.Time
	UpdatedAt time.Time
}
