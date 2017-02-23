package models

import (
	"database/sql"
	"log"

	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id    int
	Email string
}

// Create is User Data Insert DB
func (user *User) Create() error {
	db, err := sql.Open("sqlite3", "./app/test.db")
	if err != nil {
		log.Fatal(err)
	}
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into User(id, email) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id, user.Email)
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
	}
	return err
}

// Get is User get by Id
func (user *User) Get(id int) User {
	db, err := sql.Open("sqlite3", "./app/test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query(fmt.Sprintf("select id, email from User where id=%d", id))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var u User
	for rows.Next() {
		var id int
		var email string
		err = rows.Scan(&id, &email)
		if err != nil {
			log.Fatal(err)
		}
		u = User{
			Id:    id,
			Email: email,
		}
		break
	}
	return u
}

// GetAll is Get All User
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
