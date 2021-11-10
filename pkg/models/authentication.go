package models

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Auth struct {
	gorm.Model
	Email    *string
	Password string
}

// var newUser models.User

func Login(a *Auth) (*Auth, error) {

	password := []byte(a.Password)
	// hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	// if err != nil {
	// 	return nil, errors.New(err.Error())
	// }

	res := db.Where("email=?", a.Email).First(&Users)
	log.Println(res)
	if res.RowsAffected == 1 {
		if err := bcrypt.CompareHashAndPassword([]byte(a.Password), password); err != nil {
			return nil, errors.New(err.Error())
		}
	} else {
		return nil, errors.New("wrong email/password")
	}
	return a, nil
}
