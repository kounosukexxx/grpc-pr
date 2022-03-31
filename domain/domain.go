package domain

import "time"

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
