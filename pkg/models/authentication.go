package models

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
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

type Token struct {
	TokenString string
}

type TokenResponse struct {
	Role string
	Id   int64
}

func Login(a *Auth) (string, error) {

	var userinfo UserInfo
	// var token Token
	db.Table("users").Select("password", "id").Where("email = ?", a.Email).Scan(&userinfo)
	password := userinfo.Password
	// userid := userinfo.Id
	// log.Println(os.Getenv("username"))

	// validate password from user input and database
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(a.Password)); err != nil {
		return "", errors.New("wrong username/password")
	}

	token, err := GenerateToken(userinfo)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return token, nil
}

func GenerateToken(userinfo UserInfo) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = userinfo.Id
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		return "", errors.New(err.Error())
	}
	return token, nil
}
