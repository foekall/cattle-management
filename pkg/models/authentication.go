package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Auth struct {
	gorm.Model
	Email    *string
	Password string
	Token    string
}

// var newUser models.User

type Password struct {
	Password string
}

func Login(a *Auth) (*Auth, error) {

	var password Password
	db.Table("users").Select("password").Where("email = ?", a.Email).Scan(&password)
	pwd := password.Password
	if err := bcrypt.CompareHashAndPassword([]byte(pwd), []byte(a.Password)); err != nil {
		return nil, errors.New(err.Error())
	}
	return a, nil
}
