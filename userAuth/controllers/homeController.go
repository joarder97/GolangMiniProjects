package controllers

import (
	"fmt"
	"main/services"
	"net/http"
	"text/template"

	"github.com/dgrijalva/jwt-go"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tokenCookie, err := r.Cookie("token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	token, err := jwt.Parse(tokenCookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("mySigningKey"), nil
	})
	// fmt.Println(token)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if the claims have a valid "username" key
		if username, ok := claims["id"].(string); ok {
			Data, err := services.GetUserInfo(username)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			} else {
				fmt.Println(Data)
				tmpl := template.Must(template.ParseFiles("views/home.html"))
				tmpl.Execute(w, Data)
			}

		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
