package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Status   string `json:"status" form:"status"`
}

type UserResponse struct {
	ID       int    `json:"id" form:"name"`
	Username string `json:"username" form:"username"`
	Status   string `json:"status" form:"status"`
	Token    string `json:"token" form:"token"`
}
