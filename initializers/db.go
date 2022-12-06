package initializers

import (
	"log"
	"os"

	"github.com/akilans/fiber-book-rest/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect to DB
func ConnectDB() {
	var err error
	dsn := os.Getenv("DB_DSN")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to DB")
	} else {
		log.Println("Connected to DB successfully")
	}
}

// Migrate tables
func SyncDB() {
	DB.AutoMigrate(&models.Book{})
}
