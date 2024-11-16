package controllers

import (
	"net/http"

	"github.com/busyfingers/rest-api-gin-gorm/dtos"
	"github.com/busyfingers/rest-api-gin-gorm/models"
	"github.com/gin-gonic/gin"
)

// PATCH /books/:id
// Update a book
func UpdateBook(c *gin.Context) {
	var input dtos.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, ok := FetchBookOrRespond(c)
	if !ok {
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
