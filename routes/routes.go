package routes

import (
	"github.com/akilans/fiber-book-rest/controllers"
	"github.com/gofiber/fiber/v2"
)

// Define all routes and handlers call
func Routes(app *fiber.App) {

	// List all books
	app.Get("/", controllers.ListBooksHandler)

	// Add a new book
	app.Post("/addbook", controllers.AddBookHandler)

	// get a book by id
	app.Get("/book/:id", controllers.GetBookHandler)

	// update a book by id
	app.Put("/book/:id", controllers.UpdateBookHandler)

	// delete a book by id
	app.Delete("/book/:id", controllers.DeleteBookHandler)

	// Register Admin page
	app.Post("/admin", controllers.AddUserHandler)

	// Login admin
	app.Post("/login", controllers.LoginHandler)

}
