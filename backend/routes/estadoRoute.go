package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupEstadoRoutes(app *fiber.App) {
	estadoRouter := app.Group("/estados")
	estadoRouter.Get("/", controllers.ListarEstados)
	estadoRouter.Get("/:id/cidades", controllers.ListarCidadesPorEstado)
}
