package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pfirulo2022/go-blog/database"
	"github.com/pfirulo2022/go-blog/model"
)

func BlogList(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Blog List",
	}
	db := database.DBConn

	var records []model.Blog

	db.Find(&records)

	context["blog_records"] = records

	c.Status(200)
	return c.JSON(context)
}

func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Add a Blog",
	}
	c.Status(200)
	return c.JSON(context)

}

func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Update Blog",
	}
	c.Status(200)
	return c.JSON(context)
}

func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Delete Blog for the given ID",
	}
	c.Status(200)
	return c.JSON(context)
}
