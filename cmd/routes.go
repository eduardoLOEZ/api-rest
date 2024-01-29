package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/eduardoLOEZ/api-rest/handlers"
)

func setupRoutes(app *fiber.App){
	
	app.Get("/", handlers.Home)
	
	app.Get("/users", handlers.ListFacts)
	app.Post("/create", handlers.CreateUser)
}