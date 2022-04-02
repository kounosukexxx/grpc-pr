package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id    uuid.UUID
	Name  string
	Email string
	CreatedAt time.Time
	UpdatedAt time.Time
}
