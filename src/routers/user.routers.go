package routers

import (
	"github.com/PeluMike/blog/src/controllers"
	"github.com/PeluMike/blog/src/middlewares"
	"github.com/PeluMike/blog/src/utils"
	"github.com/gofiber/fiber/v2"
)

func UserRouters(api fiber.Router) {
	api.Post("/create", utils.CreateUser, controllers.CreateUser)
	api.Get("/get-all", controllers.GetAllUsers)
	api.Get("/:email", controllers.GetUserWithEmail)
	api.Get("/:email", controllers.GetUserWithEmail)
	api.Post("/login", utils.UserLogin, controllers.UserLogin)
	api.Get("/", middlewares.IsAuth, controllers.GetUser)
}
