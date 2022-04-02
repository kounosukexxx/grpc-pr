package model

import (
	"time"
)

type User struct {
	Id        string    `firestore:"id"`
	Name      string    `firestore:"name"`
	Email     string    `firestore:"email"`
	CreatedAt time.Time `firestore:"created_at"`
	UpdatedAt time.Time `firestore:"updated_at"`
}
