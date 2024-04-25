package controllers

import (
	"backend/initializers"
	"backend/models"
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

func ListBairrosByCidadeID(c *fiber.Ctx) error {
	cidadeID := c.Params("cidade_id")

	var bairros []models.Bairro
	if err := initializers.DB.Where("cidade_id = ?", cidadeID).Find(&bairros).Error; err != nil {
		return err
	}

	return c.JSON(bairros)
}
