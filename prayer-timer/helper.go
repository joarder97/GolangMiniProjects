package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func UpsertPrayerTimer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	now := time.Now()
	date := now.Format("02-01-06")

	response, err := http.Get("https://api.aladhan.com/v1/timingsByAddress/" + date + "?address=Dhaka,Bangladesh&method=8")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fajr, dhuhr, asr, maghrib, isha := processJson(data)
	updateToCurrentPrayerTime(fajr, dhuhr, asr, maghrib, isha)
	fajr, dhuhr, asr, maghrib, isha = getPrayerTimeFromDB()

	returnJson := map[string]string{"fajr": fajr, "dhuhr": dhuhr, "asr": asr, "maghrib": maghrib, "isha": isha}
	json.NewEncoder(w).Encode(returnJson)
}

func processJson(data []byte) (string, string, string, string, string) {
	var result map[string]interface{}
	json.Unmarshal([]byte(data), &result)

	fajr := result["data"].(map[string]interface{})["timings"].(map[string]interface{})["Fajr"].(string)
	dhuhr := result["data"].(map[string]interface{})["timings"].(map[string]interface{})["Dhuhr"].(string)
	asr := result["data"].(map[string]interface{})["timings"].(map[string]interface{})["Asr"].(string)
	maghrib := result["data"].(map[string]interface{})["timings"].(map[string]interface{})["Maghrib"].(string)
	isha := result["data"].(map[string]interface{})["timings"].(map[string]interface{})["Isha"].(string)

	return fajr, dhuhr, asr, maghrib, isha
}

func updateToCurrentPrayerTime(fajr, dhuhr, asr, maghrib, isha string) {
	db := ConnectDB()
	//create a row of id 1 if not exist
	_, insertError := db.Exec("INSERT INTO prayertiming (id, fajr, dhuhr, asr, maghrib, isha) VALUES (1, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE fajr = ?, dhuhr = ?, asr = ?, maghrib = ?, isha = ?", fajr, dhuhr, asr, maghrib, isha, fajr, dhuhr, asr, maghrib, isha)
	if insertError != nil {
		panic(insertError)
	}
}

func getPrayerTimeFromDB() (string, string, string, string, string) {
	db := ConnectDB()
	var fajr, dhuhr, asr, maghrib, isha string
	query := "SELECT fajr, dhuhr, asr, maghrib, isha FROM prayertiming WHERE id = 1"
	row := db.QueryRow(query)
	row.Scan(&fajr, &dhuhr, &asr, &maghrib, &isha)
	return fajr, dhuhr, asr, maghrib, isha
}

func deletePrayerTimer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := ConnectDB()
	_, deleteError := db.Exec("DELETE FROM prayertiming WHERE id = 1")
	if deleteError != nil {
		panic(deleteError)
	} else {
		returnJson := map[string]string{"message": "Prayer time deleted successfully"}
		json.NewEncoder(w).Encode(returnJson)
	}
}
