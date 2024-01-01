package main

import (
	"crud/config"
	"crud/routes"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	config.LoadConfig()
	config.ConnectDB()
}

func main() {
	app := fiber.New()
	api := app.Group("/api")

	api.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	api.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Hello World!",
		})
	})

	routes.NoteRoutes(api)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", config.ENV.PORT)))
}
