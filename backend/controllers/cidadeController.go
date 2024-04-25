package controllers

import (
	"backend/initializers"
	"backend/repositories"

	"github.com/gofiber/fiber/v2"
)

func GetAllCidades(c *fiber.Ctx) error {
	cidades, err := repositories.GetAllCidades(initializers.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao obter as cidades",
		})
	}
	return c.JSON(cidades)
}