package controllers

import (
	"context"

	"github.com/PeluMike/blog/database"
	"github.com/PeluMike/blog/src/sqlc"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateComment(c *fiber.Ctx) error {

	var commentInput sqlc.CreatePostCommentParams

	if err := c.BodyParser(&commentInput); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Something went wrong",
		})
	}
	user := c.Locals("user")
	userC, okay := user.(sqlc.GetUserByEmailRow)

	if !okay {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Something went wrong",
		})
	}

	nullId := uuid.NullUUID{UUID: userC.ID, Valid: true}

	commentInput.UserID = nullId

	comment, commErr := database.Queries.CreatePostComment(context.Background(), commentInput)

	if commErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": commErr.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Comment created",
		"data":    comment,
	})

}
