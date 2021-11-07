package controllers

import (
	"net/http"

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

	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, resAuth)

}
