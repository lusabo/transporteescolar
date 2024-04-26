package controllers

import (
	"backend/initializers"
	"backend/repositories"

	"github.com/gofiber/fiber/v2"
)

func ListarEstados(c *fiber.Ctx) error {
	estados, err := repositories.ListarEstados(initializers.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao obter os estados",
		})
	}
	return c.JSON(estados)
}
