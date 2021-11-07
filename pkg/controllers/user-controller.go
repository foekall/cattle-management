package controllers

import (
	"net/http"
	"strconv"

	"github.com/foekall/cattle-management/pkg/models"
	"github.com/gin-gonic/gin"
)

var newUser models.User

func CreateUser(c *gin.Context) {
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newUser, err := models.CreateUser(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, newUser)
}

func GetAllUser(c *gin.Context) {
	page, _ := strconv.Atoi(c.Param("page"))
	size, _ := strconv.Atoi(c.Param("size"))
	Users := models.GetAllUser(page, size)
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, Users)
}
