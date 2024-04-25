package main

import (
	"backend/initializers"
	"backend/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	initializers.LoadEnv()
	initializers.Connection()
	initializers.SyncDB()
	// initializers.Seed()
}

func main() {

	app := fiber.New()

	app.Use(cors.New())

	routes.Setup(app)

	app.Listen(":" + os.Getenv("WEB_PORT"))

}
