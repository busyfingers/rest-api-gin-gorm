package controllers

import (
	"net/http"

	"github.com/busyfingers/rest-api-gin-gorm/models"
	"github.com/gin-gonic/gin"
)

// DELETE /books/:id
// Delete a book
func DeleteBook(c *gin.Context) {
	book, ok := FetchBookOrRespond(c)
	if !ok {
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
