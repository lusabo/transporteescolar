package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupCidadeRoutes(app *fiber.App) {
	cidadeRouter := app.Group("/cidades")
	cidadeRouter.Get("/:id/bairros", controllers.ListarBairrosPorCidade)
	cidadeRouter.Get("/:id/escolas", controllers.ListarEscolasPorCidade)
}
