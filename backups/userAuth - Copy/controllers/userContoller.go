package controllers

import (
	"encoding/json"
	"main/models"
	"main/utils"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/joarder97/golang-tutorials/userAuth/config"
	"github.com/joarder97/golang-tutorials/userAuth/services"
)

// RegisterHandler handles the POST request to create a new user.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body.
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate the request data.
	if user.Name == "" || user.Email == "" || user.Password == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Name, email, and password are required fields")
		return
	}

	// Hash the password.
	hashedPassword, err := services.HashPassword(user.Password)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Error hashing password")
		return
	}

	// Set the hashed password.
	user.Password = hashedPassword

	// Set the user's creation date.
	user.CreatedAt = time.Now()

	// Generate a new MongoDB ObjectID for the user.
	user.ID = primitive.NewObjectID()

	// Insert the user into the database.
	err = services.InsertUser(user)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Error inserting user into database")
		return
	}

	// Generate a JWT token for the user.
	token, err := services.GenerateToken(user.ID.Hex(), config.JWTSecret)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Error generating JWT token")
		return
	}

	// Send the response.
	utils.SendSuccessResponse(w, http.StatusOK, map[string]string{
		"token": token,
	})
}

// LoginHandler handles the POST request to login a user.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body.
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate the request data.
	if user.Email == "" || user.Password == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Email and password are required fields")
		return
	}

	// Get the user from the database.
	dbUser, err := services.GetUserByEmail(user.Email)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Error getting user from database")
		return
	}

	// Check if the user exists.
	if dbUser.ID == "" {
		utils.SendErrorResponse(w, http.StatusUnauthorized, "Email or password is incorrect")
		return
	}

	// Compare the hashed passwords.
	if !services.ComparePasswords(dbUser.Password, user.Password) {
		utils.SendErrorResponse(w, http.StatusUnauthorized, "Email or password is incorrect")
		return
	}

	// Generate a JWT token for the user.
	token, err := services.GenerateToken(dbUser.ID.Hex(), config.JWTSecret)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Error generating JWT token")
		return
	}

	// Send the response.
	utils.SendSuccessResponse(w, http.StatusOK, map[string]string{
		"token": token,
	})
}
