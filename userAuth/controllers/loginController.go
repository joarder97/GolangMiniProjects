package controllers

import (
	"fmt"
	"main/services"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func Login(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "views/login.html")
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)

	username := r.FormValue("username")

	//check if user exists by userId
	exist, err := services.IfUserExist(username)
	fmt.Println(exist)
	if err != nil {
		http.Error(w, "Failed to check if user exists", http.StatusBadRequest)
		return
	}

	if exist {
		//hash password
		password := []byte(r.FormValue("password"))
		isValid, err := services.CheckPassword(username, string(password))
		if err != nil {
			http.Error(w, "Failed to check password", http.StatusBadRequest)
			return
		}
		if isValid {
			// Create a new JWT with the user's ID and expiry time
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"id":  username,
				"exp": time.Now().Add(time.Hour * 24).Unix(),
			})

			// Sign the token with the secret key
			tokenString, err := token.SignedString([]byte("mySigningKey"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Create a secure, http-only cookie to store the JWT
			http.SetCookie(w, &http.Cookie{
				Name:     "token",
				Value:    tokenString,
				HttpOnly: true,
				Secure:   true,
			})
			http.Redirect(w, r, "/home", http.StatusFound)
		}
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
}
