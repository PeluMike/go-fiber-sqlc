package utils

import (
	"strings"

	"github.com/PeluMike/blog/src/structs"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var body structs.UserInput
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	if err := validate.Struct(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": formatValidationError(err),
		})
	}

	c.Locals("body", body)
	return c.Next()
}

func UserLogin(c *fiber.Ctx) error {
	var body structs.LoginUser
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	if err := validate.Struct(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": formatValidationError(err),
		})
	}

	c.Locals("body", body)
	return c.Next()
}

func CreatePost(c *fiber.Ctx) error {
	var body structs.CreatePost
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Wrong data format",
		})
	}

	if validErr := validate.Struct(&body); validErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": formatValidationError(validErr),
		})
	}

	return c.Next()
}

func CreateComment(c *fiber.Ctx) error {
	var body structs.CreateComment
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Wrong data format",
		})
	}

	if validErr := validate.Struct(&body); validErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": formatValidationError(validErr),
		})
	}
	return c.Next()
}

// validator
func formatValidationError(err error) string {
	var sb strings.Builder
	for _, err := range err.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required":
			sb.WriteString(err.Field() + " is required")
		case "email":
			sb.WriteString(err.Field() + " must be a valid email address")
		case "min":
			sb.WriteString(err.Field() + " must be at least " + err.Param() + " characters long")
		// Add more cases as needed
		default:
			sb.WriteString(err.Field() + " is invalid")
		}
		sb.WriteString(", ") // Separate multiple errors
	}
	return strings.TrimSuffix(sb.String(), ", ") // Remove trailing comma and space
}
