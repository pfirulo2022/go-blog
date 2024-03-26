package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pfirulo2022/go-blog/database"
)

func main() {
	database.ConnectDB() // Connect to the database

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Bienvenido a mi web"})
	})

	app.Listen(":8000")

}
