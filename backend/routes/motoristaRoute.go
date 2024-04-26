package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupMotoristaRoutes(app *fiber.App) {
	motoristaRouter := app.Group("/motoristas")
	motoristaRouter.Get("", controllers.ListarMotoristas)
	//?bairro=:bairro_id&escola=:escola_id", nil)
	// motoristaRouter.Get("/:motorista_id", nil)
	// motoristaRouter.Post("/", nil)
}
