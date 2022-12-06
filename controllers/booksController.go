package controllers

import "github.com/gofiber/fiber/v2"

// List books function
func ListBooks(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
