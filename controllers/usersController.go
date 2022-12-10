package controllers

import (
	"log"
	"strconv"

	"github.com/akilans/fiber-book-rest/helpers"
	"github.com/akilans/fiber-book-rest/models"
	"github.com/gofiber/fiber/v2"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Add User Handler
func AddUserHandler(c *fiber.Ctx) error {
	var loginUser models.User
	log.Println(c.FormValue("password"))

	if err := c.BodyParser(&loginUser); err != nil {
		helpers.LogError(err)
		errMsg := Message{"Bad Request", "Failed to add a new user"}
		c.JSON(errMsg)
		return c.SendStatus(400)
	} else {
		hashedPassword, err := helpers.GenerateHashPassword(loginUser.Password)

		if err != nil {
			helpers.LogError(err)
			errMsg := Message{"Server Error", "Failed to Hash password"}
			c.JSON(errMsg)
			return c.SendStatus(500)
		}
		loginUser.Password = hashedPassword
		loginUserID, err := models.AddUser(loginUser)
		if err != nil {
			helpers.LogError(err)
			errMsg := Message{"Server Error", "Failed to add a new user"}
			c.JSON(errMsg)
			return c.SendStatus(500)
		} else {
			successMsg := Message{"Success", "New User added successfully with id - " + strconv.Itoa(loginUserID)}
			c.JSON(successMsg)
			return c.SendStatus(200)
		}
	}
}

// Login Handler
func LoginHandler(c *fiber.Ctx) error {
	var loginUser Login
	if err := c.BodyParser(&loginUser); err != nil {
		helpers.LogError(err)
		errMsg := Message{"Bad Request", "Failed to login"}
		c.JSON(errMsg)
		return c.SendStatus(400)
	} else {
		user, err := models.GetUserByEmail(loginUser.Email)
		if err != nil {
			helpers.LogError(err)
			errMsg := Message{"Server Error", "Failed to login a user"}
			c.JSON(errMsg)
			return c.SendStatus(500)
		} else {
			if (user == models.User{}) {
				errMsg := Message{"Login Failed", "User not found"}
				return c.JSON(errMsg)
			}
			result := helpers.CheckHashPassword(loginUser.Password, user.Password)
			if result {
				successMsg := Message{"Success", "Login successful"}
				return c.JSON(successMsg)
			} else {
				errMsg := Message{"Login Failed", "Invalid password"}
				return c.JSON(errMsg)
			}
		}
	}
}
