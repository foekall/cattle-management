package controllers

import (
	"net/http"
	"strconv"

	"github.com/foekall/cattle-management/pkg/models"
	"github.com/gin-gonic/gin"
)

// var newCattle models.Cattle

func CreateCattle(c *gin.Context) {

	// CreateCattle := models.CreateCattle()
	var input models.Cattle
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	NewCattle := models.CreateCattle(&input)
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, NewCattle)

}

func UpdateCattle(c *gin.Context) {

	var cattle models.Cattle
	cid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cattle.ID = cid

	if err := c.ShouldBindJSON(&cattle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	NewCattle, err := models.UpdateCattle(&cattle)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, NewCattle)

}

func DeleteCattle(c *gin.Context) {
	cid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Message, err := models.DeleteCattle(cid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"message": Message})
}

func GetAllCattles(c *gin.Context) {
	NewCattle := models.GetAllCattles()
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, NewCattle)
}

func GetCattleById(c *gin.Context) {
	id := c.Param("id")
	book_id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Cattle := models.GetCattleById(book_id)
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, Cattle)
}
