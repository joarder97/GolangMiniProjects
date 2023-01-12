package main

import (
	"main/config"
	"main/services"
)

func main() {
	err := config.Load()
	if err != nil {
		panic(err)
	}

	err = services.ConnectToDB(config.Cfg.MongoURI)
	if err != nil {
		panic(err)
	}
	userID := "12345"

	_, err := services.GenerateToken(userID, config.Cfg.JWTSecret)
	if err != nil {
		panic(err)
	}
}
