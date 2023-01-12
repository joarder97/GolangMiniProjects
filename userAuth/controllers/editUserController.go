package controllers

import (
	"fmt"
	"main/services"
	"net/http"
)

func EditUser(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	userId := r.Form.Get("userId")
	username := r.Form.Get("username")
	email := r.Form.Get("email")
	fmt.Println(userId, username, email)
	services.EditUser(userId, username, email)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
