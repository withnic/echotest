package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/withnic/echotest/web/app/config"
)

// UserDb is DbModel
type UserDb struct{}

// GetRecord gets User's record
func (udb *UserDb) GetRecord(id int) User {
	db, err := sql.Open(config.DbType, config.DbPath)
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
			ID:    id,
			Email: email,
			Db:    UserDb{},
		}
		break
	}
	return u
}

// GetRecords is multi records
func (udb *UserDb) GetRecords() []User {
	db, err := sql.Open(config.DbType, config.DbPath)
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
			ID:    id,
			Email: email,
		}
		users = append(users, u)
	}
	return users
}

// Insert is isert record
func (udb *UserDb) Insert(u *User) error {
	db, err := sql.Open(config.DbType, config.DbPath)
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
	_, err = stmt.Exec(u.ID, u.Email)
	if err != nil {
		log.Fatal(err)
	}
	err = tx.Commit()

	if err != nil {
		log.Fatal(err)
	}
	return err
}
