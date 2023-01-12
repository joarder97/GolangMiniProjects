package models

import (
	"time"
)

// User represents a user in the database.
type User struct {
	UserId    string    `bson:"userId`
	Name      string    `bson:"name"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	AvatarURL string    `bson:"avatar_url"`
	Gender    string    `bson:"gender"`
	CreatedAt time.Time `bson:"created_at"`
}
