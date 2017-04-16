package models

import (
	"log"
	"time"

	"github.com/go-playground/validator"
	_ "github.com/mattn/go-sqlite3"
)

// Message is member of this site.
type Message struct {
	ID        int    `db:"id"`
	Body      string `db:"body" validate:"required"`
	UserID    int    `db:"user_id"`
	CreatedAt string `db:"created_at"`
}

// Validate is message validate
func (mes *Message) Validate() error {
	validate := validator.New()
	return validate.Struct(mes)
}

// GetAll is Get All User
func (mes *Message) GetAll() []Message {
	var messages []Message
	dbmap := initDB()
	_, err := dbmap.Select(&messages, "select * from Message order by created_at desc")
	if err != nil {
		log.Fatal(err)
	}
	return messages
}

// Create is Message Data Insert DB
func (mes *Message) Create() error {
	dbmap := initDB()
	mes.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := dbmap.Insert(mes)
	return err
}

// Update is Message Data update
func (mes *Message) Update() error {
	dbmap := initDB()
	_, err := dbmap.Update(mes)
	return err
}

// Delete is Message Data Delete from DB
func (mes *Message) Delete() error {
	dbmap := initDB()
	_, err := dbmap.Delete(mes)
	return err
}
