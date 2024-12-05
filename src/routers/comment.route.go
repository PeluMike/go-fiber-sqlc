package routers

import (
	"github.com/PeluMike/blog/src/controllers"
	"github.com/PeluMike/blog/src/middlewares"
	"github.com/PeluMike/blog/src/utils"
	"github.com/gofiber/fiber/v2"
)

func CommentRoutes(api fiber.Router) {
	api.Post("/create", middlewares.IsAuth, utils.CreateComment, controllers.CreateComment)
}
