package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupMotoristaRoutes(app *fiber.App) {
	motoristaRouter := app.Group("/motoristas")
	motoristaRouter.Get("", controllers.ListarMotoristas)
	motoristaRouter.Get("/:id", controllers.BuscarMotorista)
	// motoristaRouter.Post("/", nil)
}
