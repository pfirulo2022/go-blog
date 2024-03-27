package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pfirulo2022/go-blog/database"
)

func init() {
	database.ConnectDB() // Connect to the database
}

func main() {

	sqlDb, err := database.DBConn.DB()
	if err != nil {
		panic(err)
	}

	defer sqlDb.Close()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Bienvenido a mi web"})
	})

	app.Listen(":8000")

}
