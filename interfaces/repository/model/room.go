package model

import (
	"time"
)

type Room struct {
	Id        string     `firestore:"id"`
	Name      string     `firestore:"name"`
	Users     []*User    `firestore:"users"`
	Comments  []*Comment `firestore:"comments"`
	CreatedAt time.Time  `firestore:"created_at"`
	UpdatedAt time.Time  `firestore:"updated_at"`
}

type Comment struct {
	Id        string    `firestore:"id"`
	Comment   string    `firestore:"comment"`
	User      *User     `firestore:"user"`
	CreatedAt time.Time `firestore:"created_at"`
	UpdatedAt time.Time `firestore:"updated_at"`
}
