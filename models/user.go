package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id             int       `json:"id" form:"id"`
	Email          string    `json:"email" form:"email"`
	Name           string    `json:"name" form:"name"`
	OpenPassword   string    `json:"password" form:"password"`
	HashedPassword string    `json:"-"`
	Created        time.Time `json:"created"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email, validation.Length(2, 255)),
		validation.Field(&u.OpenPassword, validation.Required, validation.Length(8, 100)),
		validation.Field(&u.Name, validation.Required, validation.Length(2, 255)),
	)
}

func (u *User) BeforeCreate() error {
	if len(u.OpenPassword) > 0 {
		h, err := hashString(u.OpenPassword)
		if err != nil {
			return err
		}
		u.HashedPassword = h
		u.OpenPassword = ""
	}
	return nil
}

func hashString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (u *User) CheckPassword(pass string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(pass)); err != nil {
		return false
	}
	return true
}
