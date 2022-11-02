package database

import (
	"erp/config"
	"erp/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetBooks() (interface{}, error) {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
