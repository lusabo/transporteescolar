package controllers

import (
	"backend/initializers"
	"backend/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ListarBairrosPorCidade(c *fiber.Ctx) error {

	cidadeId, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID da cidade inválido",
		})
	}

	bairros, err := repositories.ListarBairrosPorCidade(initializers.DB, uint(cidadeId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao obter os bairros por cidade",
		})
	}

	return c.JSON(bairros)
}
