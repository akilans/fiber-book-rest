package controllers

import (
	"strconv"

	"github.com/akilans/fiber-book-rest/helpers"
	"github.com/akilans/fiber-book-rest/models"
	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// List books function
func ListBooksHandler(c *fiber.Ctx) error {
	books, err := models.GetBooks()
	if err != nil {
		helpers.LogError(err)
		errMsg := Message{"Server Error", "Failed to list books"}
		c.JSON(errMsg)
		return c.SendStatus(500)
	} else {
		c.JSON(books)
		return c.SendStatus(200)
	}
}

// Add a new book function
func AddBookHandler(c *fiber.Ctx) error {
	var newBook models.Book

	if err := c.BodyParser(&newBook); err != nil {
		helpers.LogError(err)
		errMsg := Message{"Bad Request", "Failed to add a new book"}
		c.JSON(errMsg)
		return c.SendStatus(400)
	} else {
		newBookID, err := models.AddBook(newBook)
		if err != nil {
			helpers.LogError(err)
			errMsg := Message{"Server Error", "Failed to add a new book"}
			c.JSON(errMsg)
			return c.SendStatus(500)
		} else {
			successMsg := Message{"Success", "New book added successfully with id - " + strconv.Itoa(newBookID)}
			c.JSON(successMsg)
			return c.SendStatus(200)
		}
	}
}

// Get book by id function
func GetBookHandler(c *fiber.Ctx) error {
	bookId, err := c.ParamsInt("id")
	if err != nil {
		helpers.LogError(err)
		errMsg := Message{"Bad Request", "Provide valid book id"}
		c.JSON(errMsg)
		return c.SendStatus(400)
	}

	book := models.GetBookByID(bookId)

	if (book == models.Book{}) {
		errMsg := Message{"Not Found", "Book doesn't exists"}
		c.JSON(errMsg)
		return c.SendStatus(404)
	} else {
		c.JSON(book)
		return c.SendStatus(200)
	}

}

// Update a book by id function
func UpdateBookHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

// List books function
func DeleteBookHandler(c *fiber.Ctx) error {
	bookId, err := c.ParamsInt("id")
	if err != nil {
		helpers.LogError(err)
		errMsg := Message{"Bad Request", "Provide valid book id"}
		c.JSON(errMsg)
		return c.SendStatus(400)
	}

	err = models.DeleteBookByID(bookId)
	if err != nil {
		helpers.LogError(err)
		errMsg := Message{"Not Found", "Book doesn't exists"}
		c.JSON(errMsg)
		return c.SendStatus(404)
	} else {
		successMsg := Message{"Success", "Book deleted successfully "}
		c.JSON(successMsg)
		return c.SendStatus(200)
	}
}
