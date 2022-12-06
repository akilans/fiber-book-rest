package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akilans/fiber-book-rest/initializers"
	"github.com/gofiber/fiber/v2"
)

// Inital function to load env and connect to DB
func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
	initializers.SyncDB()
}

// Main function
func main() {
	fmt.Println("Bookstore REST API with MySQL, GORM, JWT, and Fiber")

	// setup app
	app := fiber.New()

	// router config
	Routes(app)

	PORT := os.Getenv("PORT")
	log.Println("Server started on port - ", PORT)
	// start app
	log.Fatal(app.Listen(PORT))
}
