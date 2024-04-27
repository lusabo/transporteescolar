package controllers

import (
	"backend/initializers"
	"backend/repositories"

	"github.com/gofiber/fiber/v2"
)

func ListarCidadesPorEstado(c *fiber.Ctx) error {

	id := c.Params("id")

	cidades, err := repositories.ListarCidadesPorEstado(initializers.DB, id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao obter as cidades por estado",
		})
	}

	return c.JSON(cidades)
}
