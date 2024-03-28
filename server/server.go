package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pfirulo2022/go-blog/database"
	"github.com/pfirulo2022/go-blog/router"
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

	app.Use(logger.New())

	router.SetupRoutes(app)

	app.Listen(":8000")

}
