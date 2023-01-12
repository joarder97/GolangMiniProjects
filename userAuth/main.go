package main

import (
	"main/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/createUser", controllers.CreateUser).Methods("GET")
	router.HandleFunc("/createUserPost", controllers.CreateUserPost).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("GET")
	router.HandleFunc("/loginPost", controllers.LoginPost).Methods("POST")
	router.HandleFunc("/home", controllers.Home).Methods("GET")
	router.HandleFunc("/logout", controllers.Logout).Methods("GET")
	//edit user as username param
	router.HandleFunc("/editUser", controllers.EditUser).Methods("Post")
	http.ListenAndServe(":8000", router)
}
