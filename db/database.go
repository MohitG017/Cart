package db

import (
	"backend/models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Import SQLite dialect
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// Update with your SQLite database file path
	dsn := "gorm.db" // SQLite database file
	db, err := gorm.Open("sqlite3", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate the models
	db.AutoMigrate(&models.User{}, &models.Cart{}, &models.Item{}, &models.Order{}, &models.CartItem{})

	DB = db
	return db
}
