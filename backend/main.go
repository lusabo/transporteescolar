package main

import (
	"handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Defina suas rotas aqui
	app.Get("/", handlers.Home)

	// Inicie o servidor na porta 3000
	app.Listen(":3000")
}
