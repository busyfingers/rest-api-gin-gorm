package controllers

import (
	"net/http"

	queries "github.com/busyfingers/rest-api-gin-gorm/data"
	"github.com/busyfingers/rest-api-gin-gorm/models"
	"github.com/gin-gonic/gin"
)

func FetchBookOrRespond(c *gin.Context) (*models.Book, bool) {
	// Call the common query function
	book, err := queries.GetBookById(c.Param("id"))
	if err != nil {
		// Handle the error by responding and return `false` for success
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, false
	}

	// Return the book and indicate success
	return book, true
}
