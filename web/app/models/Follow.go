package models

import validator "gopkg.in/go-playground/validator.v9"

type Follow struct {
	ID        int    `db:"id"`
	UserID    string `db:"user_id" validate:"required"`
	FollowID  string `db:"follow_id" validate:"required"`
	CreatedAt string `db:"created_at"`
}

// Validate is Follow validater
func (f *Follow) Validate() error {
	validate := validator.New()
	return validate.Struct(f)
}
