package controllers

import (
	"net/http"
	"strconv"

	"github.com/foekall/cattle-management/pkg/models"
	"github.com/foekall/cattle-management/pkg/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var newUser models.User

var validate *validator.Validate

func CreateUser(c *gin.Context) {

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate = validator.New()
	validate.RegisterValidation("is-duplicate-email", validators.ValidateEmail)
	err := validate.Struct(&newUser)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Field() + " is a required"})
				return
			case "max":
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Field() + " is shout not more than " + err.Param() + " characters"})
				return
			case "min":
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Field() + " is shout not less than " + err.Param() + " characters"})
				return
			case "email":
				c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong email address"})
				return
			case "is-duplicate-email":
				c.JSON(http.StatusBadRequest, gin.H{"error": "email duplicate"})
				return
			default:
				c.JSON(http.StatusBadRequest, gin.H{"error": "Something not right"})
				return
			}
		}
	}

	newUser, err := models.CreateUser(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newUser)
}

func GetAllUser(c *gin.Context) {
	page, _ := strconv.Atoi(c.Param("page"))
	size, _ := strconv.Atoi(c.Param("size"))
	Users := models.GetAllUser(page, size)
	c.JSON(http.StatusOK, Users)
}
