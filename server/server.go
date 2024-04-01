package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/pfirulo2022/go-blog/database"
	"github.com/pfirulo2022/go-blog/router"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB() // Connect to the database
}

func main() {

	sqlDb, err := database.DBConn.DB()
	if err != nil {
		panic(err)
	}

	defer sqlDb.Close()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "application/json; charset=utf-8",
	}))

	app.Use(logger.New())

	router.SetupRoutes(app)

	app.Listen(":8000")

}
