package controllers

import (
	"context"
	"strings"

	"github.com/PeluMike/blog/database"
	"github.com/PeluMike/blog/src/config"
	"github.com/PeluMike/blog/src/sqlc"
	"github.com/PeluMike/blog/src/structs"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	var userInput sqlc.CreateUserParams
	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(500).JSON(fiber.Map{})
	}
	_, err := database.Queries.GetUserByEmail(context.Background(), userInput.Email)

	if err == nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "User already exist",
		})
	}

	password := []byte(userInput.Password)
	hashedPass, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	createUserParams := sqlc.CreateUserParams{
		FirstName: strings.ToLower(userInput.FirstName),
		Email:     strings.ToLower(userInput.Email),
		Password:  string(hashedPass),
		LastName:  strings.ToLower(userInput.LastName),
	}
	user, err := database.Queries.CreateUser(context.Background(), createUserParams)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  true,
		"user":    user,
		"message": "User created successfully",
	})

}

// login user
func UserLogin(c *fiber.Ctx) error {
	var userInput structs.LoginUser
	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(500).JSON(fiber.Map{})
	}

	user, err := database.Queries.GetUserWithPass(context.Background(), userInput.Email)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email or password incorrect",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email or password incorrect",
		})
	}

	token, err := config.GenerateJWT(userInput.Email)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Something went wrong!",
		})
	}

	// c.Cookie(&fiber.Cookie{
	// 	Name:     "jwt",
	// 	Value:    token,
	// 	HTTPOnly: true,
	// 	Secure:   true,
	// })

	return c.Status(201).JSON(fiber.Map{
		"message": "User login successfully",
		"token":   token,
	})
}

// all users
func GetAllUsers(c *fiber.Ctx) error {
	users, err := database.Queries.GetUsers(context.Background())
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"users":   users,
		"message": "All users retrieved",
	})
}

// user with email
func GetUserWithEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	user, err := database.Queries.GetUserByEmail(context.Background(), strings.ToLower(email))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "User with email doesn't exist",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "User data retrieved successfully!",
		"data":    user,
	})
}

// user info
func GetUser(c *fiber.Ctx) error {

	user := c.Locals("user")

	return c.Status(200).JSON(fiber.Map{
		"message": "user data",
		"user":    user,
	})
}
