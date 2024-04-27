package controllers

import (
	"backend/initializers"
	"backend/repositories"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func ListarMotoristas(c *fiber.Ctx) error {

	//TODO fazer tratamento de erro
	bairroId, err := strconv.ParseUint(c.Query("bairro"), 10, 32)
	escolaId, err := strconv.ParseUint(c.Query("escola"), 10, 32)

	motoristas, err := repositories.ListarMotoristas(initializers.DB, uint(bairroId), uint(escolaId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao obter os motoristas por Bairro e/ou Escola",
		})
	}

	return c.JSON(motoristas)
}

func BuscarMotorista(c *fiber.Ctx) error {

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID do motorista inválido",
		})
	}

	motorista, err := repositories.BuscarMotorista(initializers.DB, uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Não existe motorista com esse ID",
		})
	}
	return c.JSON(motorista)
}
