package main

import (
	"database/sql"
	"fmt"
)

func ConnectDB() sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/BookingApp")
	if err != nil {
		panic(err)
	} else if err = db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()
	_, nameTableCreationError := db.Exec("CREATE TABLE IF NOT EXISTS Names (id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY, name TEXT NOT NULL)")
	if err != nil {
		panic(nameTableCreationError)
	} else {
		fmt.Printf("Ok")
	}

	_, bookingTableCreationError := db.Exec("CREATE TABLE IF NOT EXISTS Booking (id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY, count INT NOT NULL)")
	if err != nil {
		panic(bookingTableCreationError)
	} else {
		_, err = db.Exec("UPDATE Booking SET count = 10 WHERE id = 1")
		if err != nil {
			panic(bookingTableCreationError)
		} else {
			fmt.Printf("Ok")
		}
	}

	return *db
}
