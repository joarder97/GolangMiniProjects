package services

import (
	"context"
	"main/models"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func ConnectToDB(uri string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	return nil
}

// InsertUser inserts a user into the database.
func InsertUser(user models.User) error {
	collection := client.Database("yourdatabase").Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}
