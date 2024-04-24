package main

import (
	"backend/controllers"
	"backend/initializers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	initializers.LoadEnv()
	initializers.Connection()
	initializers.SyncDB()
}

func main() {

	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", controllers.HelloIndex)
	app.Get("/json", controllers.HelloJson)

	app.Listen(":" + os.Getenv("WEB_PORT"))

}
