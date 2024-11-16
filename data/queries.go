package queries

import "github.com/busyfingers/rest-api-gin-gorm/models"

func GetBookById(id string) (*models.Book, error) {
	var book models.Book

	if err := models.DB.Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}

	return &book, nil
}
