package models

import _ "github.com/mattn/go-sqlite3"

// User is member of this site.
type User struct {
	ID    int
	Email string
	Db    UserDb
}

// Create is User Data Insert DB
func (user *User) Create() error {
	return user.Db.Insert(user)
}

// Get is User get by Id
func (user *User) Get(id int) User {
	return user.Db.GetRecord(id)
}

// GetAll is Get All User
func (user *User) GetAll() []User {
	return user.Db.GetRecords()
}
