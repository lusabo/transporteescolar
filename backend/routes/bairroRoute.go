package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupBairroRoutes(app *fiber.App) {
	bairroRouter := app.Group("/bairros")
	bairroRouter.Post("/", controllers.CreateBairro)
}
