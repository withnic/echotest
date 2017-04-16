package models

import (
	"time"

	"github.com/go-playground/validator"
)

type Follow struct {
	ID        int    `db:"id"`
	UserID    int    `db:"user_id" validate:"required"`
	FollowID  int    `db:"follow_id" validate:"required"`
	CreatedAt string `db:"created_at"`
}

// Validate is Follow validater
func (f *Follow) Validate() error {
	validate := validator.New()
	return validate.Struct(f)
}

// Create is User Data Insert DB
func (f *Follow) Create() error {
	dbmap := initDB()
	f.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	err := dbmap.Insert(f)
	return err
}

// DeleteByUserFollowID is Get Follow model
func (f *Follow) DeleteByUserFollowID() error {
	dbmap := initDB()
	_, err := dbmap.Exec("delete  from Follow where user_id=? and follow_id = ?", f.UserID, f.FollowID)
	return err
}

// GetByUserFollowID is Get Follow model
func (f *Follow) GetByUserFollowID() error {
	dbmap := initDB()
	err := dbmap.SelectOne(&f, "select * from Follow where user_id=? and follow_id = ?", f.UserID, f.FollowID)
	return err
}

// Delete is Follow Data Delete from DB
func (f *Follow) Delete() error {
	dbmap := initDB()
	_, err := dbmap.Delete(f)
	return err
}
