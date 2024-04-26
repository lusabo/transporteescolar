package controllers

import (
	"backend/initializers"
	"backend/repositories"

	"github.com/gofiber/fiber/v2"
)

func ListarCidadesPorEstado(c *fiber.Ctx) error {

	uf := c.Params("uf")

	if uf == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "O parâmetro UF é obrigatório",
		})
	}

	cidades, err := repositories.ListarCidadesPorEstado(initializers.DB, uf)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao obter as cidades por UF",
		})
	}

	return c.JSON(cidades)
}
