package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in the database.
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	AvatarURL string             `bson:"avatar_url"`
	Gender    string             `bson:"gender"`
	CreatedAt time.Time          `bson:"created_at"`
}
