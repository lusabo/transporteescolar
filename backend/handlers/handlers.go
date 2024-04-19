package handlers

import "github.com/gofiber/fiber/v2"

// Home é o manipulador para a rota "/"
func Home(c *fiber.Ctx) error {
	return c.SendString("Olá, mundo!")
}
