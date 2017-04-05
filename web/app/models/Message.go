package models

import (
	"database/sql"

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

// Create is Message Data Insert DB
func (mes *Message) Create() error {
	dbmap := initMessageDB()
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
