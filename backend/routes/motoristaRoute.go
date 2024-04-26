package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupMotoristaRoutes(app *fiber.App) {
	motoristaRouter := app.Group("/motoristas")
	motoristaRouter.Get("/", nil)
	motoristaRouter.Get("/:motorista_id", nil)
	motoristaRouter.Post("/", nil)
}
