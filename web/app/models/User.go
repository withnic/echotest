package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id    int
	Email string
}

func (user *User) GetAll() []User {
	db, err := sql.Open("sqlite3", "./app/test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("select id, email from User")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var users []User

	for rows.Next() {
		var id int
		var email string
		err = rows.Scan(&id, &email)
		if err != nil {
			log.Fatal(err)
		}
		u := User{
			Id:    id,
			Email: email,
		}
		users = append(users, u)
	}

	return users
}
