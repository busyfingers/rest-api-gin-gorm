package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /books/:id
// Find a book
func FindBook(c *gin.Context) { // Get model if exist

	book, ok := FetchBookOrRespond(c)
	if !ok {
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})

}
