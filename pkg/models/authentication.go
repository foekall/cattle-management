package models

import (
	"errors"

	"gorm.io/gorm"
)

type Auth struct {
	gorm.Model
	Email    *string
	Password string
}

func Login(a *Auth) (*Auth, error) {
	var Auths []Auth
	res := db.Where("email=? AND password=?", a.Email, a.Password).Find(&Auths)
	if res.RowsAffected != 1 {
		return nil, errors.New("Wrong email/password")
	}
	return a, nil
}
