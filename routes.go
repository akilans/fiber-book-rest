package main

import (
	"github.com/akilans/fiber-book-rest/controllers"
	"github.com/gofiber/fiber/v2"
)

// Define all routes and handlers call
func Routes(app *fiber.App) {
	app.Get("/", controllers.ListBooks)
}
