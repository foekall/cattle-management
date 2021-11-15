package models

import (
	"errors"
	"log"
	"os"

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

type UserInfo struct {
	Id       int64
	Password string
}

func Login(a *Auth) (*Auth, error) {

	var userinfo UserInfo
	db.Table("users").Select("password", "id").Where("email = ?", a.Email).Scan(&userinfo)
	password := userinfo.Password
	// userid := userinfo.Id
	log.Println(os.Getenv("username"))

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(a.Password)); err != nil {
		return nil, errors.New("wrong username/password")
	}
	return a, nil
}
