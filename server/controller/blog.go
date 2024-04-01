package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pfirulo2022/go-blog/database"
	"github.com/pfirulo2022/go-blog/model"
)

func BlogDetail(c *fiber.Ctx) error {
	c.Status(400)
	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not Found.")
		context["msg"] = "Record not Found."

		c.Status(404)
		return c.JSON(context)
	}

	context["record"] = record
	context["statusText"] = "Ok"
	context["msg"] = "Blog Detail"
	c.Status(200)
	return c.JSON(context)

}

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

	record := new(model.Blog)
	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request.")
		context["statusText"] = ""
		context["msg"] = "Algo salio mal."
	}

	result := database.DBConn.Create(record)
	if result.Error != nil {
		log.Println("Error al guardar los datos.")
		context["statusText"] = ""
		context["msg"] = "Hubo un error."
	}

	context["msg"] = "Guardado correctamente."
	context["data"] = record

	c.Status(201)
	return c.JSON(context)

}

func BlogUpdate(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Update Blog",
	}

	//http://localhost:8000/2

	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not Found.")

		context["statusText"] = ""
		context["msg"] = "Record not Found."
		c.Status(400)
		return c.JSON(context)
	}

	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request.")

		context["msg"] = "Something went wrong."
		c.Status(400)
		return c.JSON(context)
	}

	result := database.DBConn.Save(record)

	if result.Error != nil {
		log.Println("Error in saving data.")

		context["msg"] = "Error in saving data."
		c.Status(400)
		return c.JSON(context)
	}

	context["msg"] = "Record updated successfully."
	context["data"] = record

	c.Status(200)
	return c.JSON(context)
}

func BlogDelete(c *fiber.Ctx) error {

	c.Status(400)

	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	id := c.Params("id")

	var record model.Blog

	database.DBConn.First(&record, id)

	if record.ID == 0 {
		log.Println("No se encontro el registro.")
		context["msg"] = "No blog found with the provided ID."
		return c.JSON(context)
	}
	result := database.DBConn.Delete(record)
	if result.Error != nil {
		context["msg"] = "Algo salio mal."
		return c.JSON(context)
	}

	context["statusText"] = "ok"
	context["msg"] = "Record deleted successfully."

	c.Status(200)
	return c.JSON(context)
}
