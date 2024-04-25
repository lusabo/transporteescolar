package controllers

import (
	"backend/initializers"
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

func CreateBairro(c *fiber.Ctx) error {
	bairro := new(models.Bairro)

	if err := c.BodyParser(bairro); err != nil {
		return err
	}

	if err := initializers.DB.Create(bairro).Error; err != nil {
		return err
	}

	return c.JSON(bairro)
}
