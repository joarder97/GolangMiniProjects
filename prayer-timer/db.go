package main

import (
	"database/sql"
)

func ConnectDB() sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/prayertimingapp")
	if err != nil {
		panic(err)
	} else if err = db.Ping(); err != nil {
		panic(err)
	} else {
		//create a table named prayertiming if not exist
		_, createTableError := db.Exec("CREATE TABLE IF NOT EXISTS prayertiming (id int(11) NOT NULL AUTO_INCREMENT, fajr varchar(255) NOT NULL, dhuhr varchar(255) NOT NULL, asr varchar(255) NOT NULL, maghrib varchar(255) NOT NULL, isha varchar(255) NOT NULL, PRIMARY KEY (id)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;")
		if createTableError != nil {
			panic(createTableError)
		}
		println("Database connected successfully")
	}
	return *db
}
