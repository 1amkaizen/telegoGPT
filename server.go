package main

import (
	"telegoGPT/controllers/servercontroller"
	"telegoGPT/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func Server() {
	models.ConnectDatabase()
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	api := app.Group("/api") // /api
	user := api.Group("/users")

	user.Get("/", servercontroller.Index)
	user.Get("/:id", servercontroller.Show)
	user.Post("/", servercontroller.Create)
	user.Put("/:id", servercontroller.Update)
	user.Delete("/:id", servercontroller.Delete)

	app.Listen(":3000")

}
