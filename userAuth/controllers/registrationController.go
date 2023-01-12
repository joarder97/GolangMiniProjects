package controllers

import (
	"fmt"
	"io"
	"main/models"
	"main/services"
	"os"
	"path/filepath"
	"strings"
	"time"

	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "views/registration.html")
}

// post function for CreateUser
func CreateUserPost(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	userId := strings.Replace(r.FormValue("username"), " ", "", -1)
	// fmt.Println(userId)

	//check if user exists by userId
	exist, err := services.IfUserExist(userId)
	fmt.Println(exist)
	if err != nil {
		http.Error(w, "Failed to check if user exists", http.StatusBadRequest)
		return
	}
	if exist {
		//return user exists view
		http.ServeFile(w, r, "views/userExists.html")
	}

	file, header, err := r.FormFile("profile_picture")
	if err != nil {
		http.Error(w, "Failed to get file from form", http.StatusBadRequest)
		return
	}

	savePath := filepath.Join(".", "images")
	if _, err := os.Stat(savePath); os.IsNotExist(err) {
		os.MkdirAll(savePath, os.ModePerm)
	}
	fileName := filepath.Join(savePath, header.Filename)

	f, err := os.Create(fileName)
	if err != nil {
		http.Error(w, "Failed to create file", http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	//hash the password
	password := []byte(r.FormValue("password"))
	hashedPassword, err := bcrypt.GenerateFromPassword(password, 14)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	user := models.User{}
	user.UserId = userId
	user.Name = r.FormValue("username")
	user.Email = r.FormValue("email")
	user.Password = string(hashedPassword)
	user.AvatarURL = "images/" + header.Filename
	user.Gender = r.FormValue("gender")
	user.CreatedAt = time.Now()

	services.ConnectToDb()

	services.CreateUser(user)
	http.ServeFile(w, r, "views/login.html")
}
