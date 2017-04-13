package models

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	validator "gopkg.in/go-playground/validator.v9"

	_ "github.com/mattn/go-sqlite3"
)

// User is member of this site.
type User struct {
	ID     int    `db:"id"`
	Email  string `db:"email" validate:"required,email"`
	Passwd string `db:"passwd" validate:"required"`
	//	VPasswd string `validate:"required"`
}

// Validate is User validater
func (user *User) Validate() error {
	validate := validator.New()
	return validate.Struct(user)
}

// convertPass is Password exec bcrypt
func (user *User) generatePasswd() string {
	password := []byte(user.Passwd)
	cost := 10
	hash, _ := bcrypt.GenerateFromPassword(password, cost)
	return string(hash)
}

// Create is User Data Insert DB
func (user *User) Create() error {
	user.Passwd = user.generatePasswd()

	dbmap := initDB()
	err := dbmap.Insert(user)
	return err
}

// Update is User Data update
func (user *User) Update() error {
	user.Passwd = user.generatePasswd()
	dbmap := initDB()
	_, err := dbmap.Update(user)
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

// GetByEmail is Get User model
func (user *User) GetByEmail() error {
	dbmap := initDB()
	err := dbmap.SelectOne(&user, "select * from User where email=?", user.Email)
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

// Messages is Get Related All messages
func (user *User) Messages() []Message {
	var mes []Message
	dbmap := initDB()
	_, err := dbmap.Select(&mes, "select * from Message where user_id = :id order by created_at desc", map[string]interface{}{
		"id": user.ID,
	})
	if err != nil {
		log.Fatal(err)
	}
	return mes
}

func (user *User) Followers() []User {
	var users []User
	dbmap := initDB()
	_, err := dbmap.Select(&users, "select User.id, User.email, User.passwd from Follow inner join User on user.id = Follow.follow_id  where Follow.user_id = :id order by Follow.created_at desc", map[string]interface{}{
		"id": user.ID,
	})
	if err != nil {
		log.Fatal(err)
	}

	return users
}

// GetAllFollowersAndMeMessage is Get All Follow message
func (user *User) GetAllFollowersAndMeMessage() []MessageWithUser {
	var messages []MessageWithUser
	dbmap := initDB()
	rows, err := dbmap.Query("select Message.id, Message.body, Message.user_id, Message.created_at, User.id, User.email, User.passwd from Message inner join User on Message.user_id = User.id left join Follow ON Message.user_id = Follow.follow_id  where Follow.user_id = ? or User.id = ? order by Message.created_at desc", user.ID, user.ID)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var mes Message
		var user User
		rows.Scan(
			&mes.ID,
			&mes.Body,
			&mes.UserID,
			&mes.CreatedAt,
			&user.ID,
			&user.Email,
			&user.Passwd,
		)
		fmt.Print(mes)
		messages = append(messages, MessageWithUser{
			Mes:  mes,
			User: user,
		})
	}
	return messages
}
