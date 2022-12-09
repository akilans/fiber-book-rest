package controllers

import (
	"strconv"

	"github.com/akilans/fiber-book-rest/helpers"
	"github.com/akilans/fiber-book-rest/models"
	"github.com/gofiber/fiber/v2"
)

// Add User Handler
func AddUserHandler(c *fiber.Ctx) error {
	var newUser models.User
	plainPassword := c.FormValue("password")
	hashedPassword, err := helpers.GenerateHashPassword(plainPassword)

	if err != nil {
		helpers.LogError(err)
		errMsg := Message{"Server Error", "Failed to Hash password"}
		c.JSON(errMsg)
		return c.SendStatus(500)
	}

	if err := c.BodyParser(&newUser); err != nil {
		helpers.LogError(err)
		errMsg := Message{"Bad Request", "Failed to add a new user"}
		c.JSON(errMsg)
		return c.SendStatus(400)
	} else {
		newUser.Password = hashedPassword
		newUserID, err := models.AddUser(newUser)
		if err != nil {
			helpers.LogError(err)
			errMsg := Message{"Server Error", "Failed to add a new user"}
			c.JSON(errMsg)
			return c.SendStatus(500)
		} else {
			successMsg := Message{"Success", "New User added successfully with id - " + strconv.Itoa(newUserID)}
			c.JSON(successMsg)
			return c.SendStatus(200)
		}
	}
}
