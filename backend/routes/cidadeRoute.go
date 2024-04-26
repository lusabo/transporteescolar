package routes

import (
	"backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupCidadeRoutes(app *fiber.App) {
	cidadeRouter := app.Group("/cidades")
	cidadeRouter.Get("/:cidade_id/bairros", controllers.ListBairrosByCidadeID)
	cidadeRouter.Get("/:cidade_id/escolas", controllers.ListEscolasByCidadeID)
}
