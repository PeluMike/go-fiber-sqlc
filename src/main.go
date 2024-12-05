package main

import (
	"github.com/PeluMike/blog/database"
	"github.com/PeluMike/blog/src/routers"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func init() {
	database.ConnectDb()
}

func main() {
	app := fiber.New()

	api := app.Group("/api/user")
	postRoutes := app.Group("/api/posts")
	commentRoutes := app.Group("/api/comments")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Welocme to this api",
		})
	})

	routers.UserRouters(api)
	routers.PostRoutes(postRoutes)
	routers.CommentRoutes(commentRoutes)

	app.Listen(":4000")
}
