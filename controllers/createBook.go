package controllers

import (
	"net/http"

	"github.com/busyfingers/rest-api-gin-gorm/dtos"
	"github.com/busyfingers/rest-api-gin-gorm/models"
	"github.com/gin-gonic/gin"
)

// POST /books
// Create new book
func CreateBook(c *gin.Context) {
	// Validate input
	var input dtos.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
