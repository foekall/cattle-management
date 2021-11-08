package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/foekall/cattle-management/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var newUser models.User
var validate *validator.Validate

func CreateUser(c *gin.Context) {
	if err := c.ShouldBindJSON(&newUser); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate = validator.New()
	err := validate.Struct(&newUser)
	if err != nil {

		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Field() + " is a required"})
			case "max":
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Field() + " is shout not more than " + err.Param() + " characters"})
			case "min":
				c.JSON(http.StatusBadRequest, gin.H{"error": "Something not right"})
			default:
			}
			// fmt.Println(err.Error())
			// fmt.Println(err.Namespace())
			// fmt.Println(err.Field())
			// fmt.Println(err.StructNamespace())
			// fmt.Println(err.StructField())
			// fmt.Println(err.Tag())
			// fmt.Println(err.ActualTag())
			// fmt.Println(err.Kind())
			// fmt.Println(err.Type())
			// fmt.Println(err.Value())
			// fmt.Println(err.Param())
			// fmt.Println()
			return
		}
	}

	newUser, err := models.CreateUser(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, newUser)
}

func GetAllUser(c *gin.Context) {
	page, _ := strconv.Atoi(c.Param("page"))
	size, _ := strconv.Atoi(c.Param("size"))
	Users := models.GetAllUser(page, size)
	// c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, Users)
}
