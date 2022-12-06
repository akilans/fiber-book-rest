package models

import (
	"log"

	"github.com/akilans/fiber-book-rest/initializers"
	"gorm.io/gorm"
)

// Book Type -> Books table
type Book struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Author   string  `json:"author"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"image_url"`
}

var db *gorm.DB

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
	db = initializers.GetDB()
	SyncDB()
}

// add a book
// Get books
func GetBooks() ([]Book, error) {
	var books []Book
	result := db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

// add a book
func AddBook(book Book) (id int, err error) {
	result := db.Create(&book)
	if result.Error != nil {
		return 0, result.Error
	} else {
		return book.ID, nil
	}
}

// Migrate tables
func SyncDB() {
	log.Println("Start of DB migration")
	err := db.AutoMigrate(&Book{})
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("End of DB migration")
}
