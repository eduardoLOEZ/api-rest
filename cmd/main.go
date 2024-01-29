package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/eduardoLOEZ/api-rest/database"

)


func main(){
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, this is a system for users")
	})

	app.Listen(":3000")
}