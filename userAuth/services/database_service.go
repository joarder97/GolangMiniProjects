package services

import (
	"context"
	"fmt"
	"log"
	"main/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func ConnectToDb() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin-joarder97:enLK3WL0LQAUptuI@cluster0.m8qqnft.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		panic(err)
	} else {
		println("Connected to database")
	}

	//create a user table if not exist
	dbService := DatabaseService{
		Client: client,
		Name:   "test",
	}
	if err := dbService.CreateCollectionIfNotExist("users"); err != nil {
		log.Fatalf("Error creating users collection: %v", err)
	}

}

type DatabaseService struct {
	Client *mongo.Client
	Name   string
}

func (db *DatabaseService) CreateCollectionIfNotExist(collectionName string) error {
	collections, err := db.Client.Database(db.Name).ListCollectionNames(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	var exists bool
	for _, c := range collections {
		if c == collectionName {
			exists = true
			break
		}
	}

	if !exists {
		err = db.Client.Database(db.Name).CreateCollection(context.Background(), collectionName)
		if err != nil {
			return err
		} else {
			println("Collection created")
		}
	}
	return nil

}

func CreateUser(user models.User) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin-joarder97:enLK3WL0LQAUptuI@cluster0.m8qqnft.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		panic(err)
	} else {
		println("Connected to database")
	}
	dbService := DatabaseService{
		Client: client,
		Name:   "test",
	}
	collection := dbService.Client.Database("test").Collection("users")
	insertResult, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func IfUserExist(userId string) (bool, error) {
	//find user by userId
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin-joarder97:enLK3WL0LQAUptuI@cluster0.m8qqnft.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		panic(err)
	} else {
		println("Connected to database")
	}
	dbService := DatabaseService{
		Client: client,
		Name:   "test",
	}
	// var foundUser bson.M

	filter := bson.M{"userid": userId}
	count, err := dbService.Client.Database("test").Collection("users").CountDocuments(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	fmt.Println("Count: ", count)
	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func CheckPassword(userId string, password string) (bool, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin-joarder97:enLK3WL0LQAUptuI@cluster0.m8qqnft.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		panic(err)
	} else {
		println("Connected to database")
	}
	dbService := DatabaseService{
		Client: client,
		Name:   "test",
	}

	filter := bson.M{"userid": userId}
	collection := dbService.Client.Database("test").Collection("users")

	var result models.User
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	hashedPassword := result.Password

	error := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if error != nil {
		fmt.Println("Invalid password")
	} else {
		fmt.Println("Valid password")
		return true, nil
	}
	return false, nil
}

func GetUserInfo(userId string) (models.User, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin-joarder97:enLK3WL0LQAUptuI@cluster0.m8qqnft.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		panic(err)
	} else {
		println("Connected to database")
	}
	dbService := DatabaseService{
		Client: client,
		Name:   "test",
	}

	filter := bson.M{"userid": userId}
	collection := dbService.Client.Database("test").Collection("users").FindOne(context.Background(), filter)
	var result models.User
	err = collection.Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result, nil
}

func EditUser(userId, username, email string) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin-joarder97:enLK3WL0LQAUptuI@cluster0.m8qqnft.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		panic(err)
	}
	// dbService := DatabaseService{
	// 	Client: client,
	// 	Name:   "test",
	// }
	fmt.Println("Updated id", userId, "updated name:", username, email)
	// collection := dbService.Client.Database("test").Collection("users")
	// filter := bson.M{"userid": userid}
	// //find and update
	// update := bson.M{"$set": bson.M{"username": username, "email": email}}
	// _, err = collection.UpdateOne(context.Background(), filter, update)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
