package models

import (
	"database/sql"
	"log"
	"time"

	validator "gopkg.in/go-playground/validator.v9"
	gorp "gopkg.in/gorp.v1"

	_ "github.com/mattn/go-sqlite3"
	"github.com/withnic/echotest/web/app/config"
)

// Message is member of this site.
type Message struct {
	ID        int    `db:"id"`
	Body      string `form:"body" db:"body" validate:"required"`
	UserID    int    `db:"user_id"`
	CreatedAt string `db:"created_at"`
}

func (mes *Message) Validate() error {
	validate := validator.New()
	return validate.Struct(mes)
}

func initMessageDB() *gorp.DbMap {
	db, err := sql.Open(config.DbType, config.DbPath)
	if err != nil {
		panic(err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(Message{}, "Message").SetKeys(true, "id")
	return dbmap
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
	dbmap := initMessageDB()
	mes.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	err := dbmap.Insert(mes)
	return err
}

// Update is Message Data update
func (mes *Message) Update() error {
	dbmap := initMessageDB()
	_, err := dbmap.Update(mes)
	return err
}

// Delete is Message Data Delete from DB
func (mes *Message) Delete() error {
	dbmap := initMessageDB()
	_, err := dbmap.Delete(mes)
	return err
}
