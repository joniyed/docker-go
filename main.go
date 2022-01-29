package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-hrms/src/database"
)

func main() {

	database.Connect()
	database.AutoMigrate()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Base API Url")
	})

	app.Listen(":8080")

}
