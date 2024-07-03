package database

import (
	"main/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("blackshaman:10203150112Zaa!@/go_admin"))

	if err != nil {
		panic("Couln not connect to the database")
	}

	DB = database

	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{}, &models.Order{}, &models.OrderItem{})
}
