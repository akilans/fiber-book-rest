package controllers

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/akilans/fiber-book-rest/helpers"
	"github.com/akilans/fiber-book-rest/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type Login struct {
	Email    string `json:"email" validate:"required,email,min=6,max=100"`
	Password string `json:"password" validate:"required,min=6,max=15"`
}

// Custom claims needed for generating JWT token
type MyCustomClaims struct {
	UserEmail    string
	LoggedInTime string
	jwt.RegisteredClaims
}

// validate user payload
func ValidateUserStruct(user models.User) []ErrorResponse {
	var errors []ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, element)
		}
	}
	return errors
}

func ValidateLoginUserStruct(user Login) []ErrorResponse {
	var errors []ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, element)
		}
	}
	return errors
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
		errors := ValidateUserStruct(loginUser)
		if errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors)
		}
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
		errors := ValidateLoginUserStruct(loginUser)
		if errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors)
		}
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
				token, _ := CreateJWT(loginUser.Email)
				successMsg := Message{"Success", token}
				return c.JSON(successMsg)
			} else {
				errMsg := Message{"Login Failed", "Invalid password"}
				return c.JSON(errMsg)
			}
		}
	}
}

// Create JWT token
// Function to create JWT token
func CreateJWT(userEmail string) (string, error) {
	currentTime := time.Now().Format("02-01-2006 15:04:05")

	// Storing user name and loggedin time
	// Token expires in 1 hour.
	claims := MyCustomClaims{
		userEmail,
		currentTime,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			Issuer:    "Akilan",
		},
	}

	// Generate token with HS256 algorithm and custom claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with our secret key
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	return signedToken, err
}
