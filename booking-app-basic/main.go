package main

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string
}

func main() {
	db := ConnectDB()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var count int
		err := db.QueryRow("SELECT count FROM Booking WHERE id = 1").Scan(&count)
		if err != nil {
			panic(err)
		}
		if count <= 0 {
			http.ServeFile(w, r, "./templates/soldOut.ejs")
			return
		} else {
			data := map[string]int{
				"remainingTickets": count,
			}

			tmpl, err := template.ParseFiles("./templates/index.ejs")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = tmpl.Execute(w, data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	})

	http.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			name := r.FormValue("name")
			if name == "" {
				http.ServeFile(w, r, "./templates/emptyName.ejs")
				return
			}
			var count int
			err := db.QueryRow("SELECT COUNT(*) FROM names WHERE name = ?", name).Scan(&count)
			if err != nil {
				panic(err)
			} else if count > 0 {
				http.ServeFile(w, r, "./templates/nameExist.ejs")
				fmt.Printf("%d", count)
			} else {
				_, err := db.Exec("INSERT INTO names (name) VALUES (?)", name)
				if err != nil {
					panic(err)
				} else {
					_, err = db.Exec("UPDATE Booking SET count = count - 1 WHERE id = 1")
					if err != nil {
						panic(err)
					} else {
						fmt.Printf("Ok")
					}
				}
				http.ServeFile(w, r, "./templates/success.ejs")
				return
			}
		}
	})

	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		//get all the names from names table
		rows, err := db.Query("SELECT name FROM names")
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		users := []User{}

		for rows.Next() {
			var user User
			err = rows.Scan(&user.Name)
			if err != nil {
				panic(err)
			}
			users = append(users, user)
		}
		err = rows.Err()
		if err != nil {
			panic(err)
		}

		tmpl := template.Must(template.ParseFiles("./templates/list.ejs"))
		if err := tmpl.Execute(w, users); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}
