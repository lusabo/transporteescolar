package main

import (
	"backend/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	// Defina suas rotas aqui
	app.Get("/", handlers.Home)

	// Inicie o servidor na porta 3000
	app.Listen(":3000")
}
