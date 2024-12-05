package routers

import (
	"github.com/PeluMike/blog/src/controllers"
	"github.com/PeluMike/blog/src/middlewares"
	"github.com/PeluMike/blog/src/utils"
	"github.com/gofiber/fiber/v2"
)

func PostRoutes(postR fiber.Router) {
	postR.Post("/create", middlewares.IsAuth, utils.CreatePost, controllers.CreatePost)
	postR.Get("/user", middlewares.IsAuth, controllers.GetUserPosts)
}
