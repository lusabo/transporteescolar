package routes

import "github.com/gofiber/fiber/v2"

func Setup(app *fiber.App) {
	SetupEstadoRoutes(app)
	SetupCidadeRoutes(app)
	//SetupMotoristaRoutes(app)
}
