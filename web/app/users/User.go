package user

import (
	"database/sql"
	"log"

	validator "gopkg.in/go-playground/validator.v9"
	gorp "gopkg.in/gorp.v1"

	_ "github.com/mattn/go-sqlite3"
	"github.com/withnic/echotest/web/app/config"
)

// User is member of this site.
type User struct {
	ID    int    `db:"id"`
	Email string `form:"email" db:"email" validate:"required,email"`
}

func (user *User) Validate() error {
	validate := validator.New()
	return validate.Struct(user)
}

func initDB() *gorp.DbMap {
	db, err := sql.Open(config.DbType, config.DbPath)
	if err != nil {
		panic(err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(User{}, "User").SetKeys(true, "id")
	return dbmap
}

// Create is User Data Insert DB
func (user *User) Create() error {
	dbmap := initDB()
	err := dbmap.Insert(user)
	return err
}

// Delete is User Data Delete from DB
func (user *User) Delete() error {
	dbmap := initDB()
	_, err := dbmap.Delete(user)
	return err
}

// Get is User get by Id
func (user *User) Get() error {
	dbmap := initDB()
	err := dbmap.SelectOne(&user, "select * from User where id=?", user.ID)
	return err
}

// GetAll is Get All User
func (user *User) GetAll() []User {
	var users []User
	dbmap := initDB()
	_, err := dbmap.Select(&users, "select * from User order by id")
	if err != nil {
		log.Fatal(err)
	}
	return users
}
