package controllers

import (
	"context"

	"github.com/PeluMike/blog/database"
	"github.com/PeluMike/blog/src/sqlc"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	var postInput sqlc.CreatePostParams

	if err := c.BodyParser(&postInput); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "something went wrong",
		})
	}

	user := c.Locals("user")
	userC, okay := user.(sqlc.GetUserByEmailRow)

	if !okay {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user from context",
		})
	}

	postInput.UserID = userC.ID

	post, postErr := database.Queries.CreatePost(context.Background(), postInput)

	if postErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": postErr.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Post created successfull",
		"post":    post,
	})
}

func GetUserPosts(c *fiber.Ctx) error {

	user := c.Locals("user")
	userC, okay := user.(sqlc.GetUserByEmailRow)

	if !okay {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user from context",
		})
	}

	posts, err := database.Queries.GetPostsByUserID(context.Background(), userC.ID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User posts retrieved",
		"post":    posts,
	})
}
