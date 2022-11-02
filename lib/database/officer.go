package database

import (
	"erp/config"
	"erp/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetOfficers() (interface{}, error) {
	var officers []models.Officer

	if err := config.DB.Find(&officers).Error; err != nil {
		return nil, err
	}
	return officers, nil
}
