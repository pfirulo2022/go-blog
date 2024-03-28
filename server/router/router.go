package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pfirulo2022/go-blog/controller"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", controller.BlogList)
	app.Post("/", controller.BlogCreate)
	app.Put("/", controller.BlogUpdate)
	app.Delete("/", controller.BlogDelete)
}
