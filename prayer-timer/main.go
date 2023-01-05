package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/createPrayer", UpsertPrayerTimer).Methods("GET")
	router.HandleFunc("/deletePrayer", deletePrayerTimer).Methods("GET")
	router.HandleFunc("/updatePrayer", UpsertPrayerTimer).Methods("GET")
	http.ListenAndServe(":8080", router)
}
