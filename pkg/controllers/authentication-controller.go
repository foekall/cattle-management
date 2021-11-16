package controllers

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/foekall/cattle-management/pkg/models"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {

	var newAuth models.Auth

	if err := c.ShouldBindJSON(&newAuth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resAuth, err := models.Login(&newAuth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": resAuth})

}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.Request
		tokenString := strings.TrimPrefix(h.Header.Get("Authorization"), "Bearer ")
		token, err := VerifyToken(tokenString)
		if err != nil {
			log.Println(err.Error())
		}
		// log.Println(token)
		claims, ok := ExtractToken(token)
		if ok {
			log.Println(claims)
		}
		c.Next()
	}
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("secret")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractToken(token *jwt.Token) (jwt.MapClaims, bool) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		return nil, false
	}
}
