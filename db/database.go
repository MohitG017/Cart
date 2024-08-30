package db

import (
	"backend/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func InitDB() *gorm.DB {

	dsn := "gorm.db"
	db, err := gorm.Open("sqlite3", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	
	db.AutoMigrate(&models.User{}, &models.Cart{}, &models.Item{}, &models.Order{}, &models.CartItem{})

	DB = db
	return db
}
