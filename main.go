package main

import (
	"log"

	"github.com/CemHarput/bookStoreGo/config"
	"github.com/CemHarput/bookStoreGo/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)
func main() {
	app := fiber.New()
	
	app.Use(cors.New())

	config.Connect()

	api := app.Group("/api/v1")

    api.Get("/books", handlers.GetBooks)
    api.Get("/books/:id", handlers.GetBook)
    api.Post("/books", handlers.AddBook)
    api.Put("/books/:id", handlers.UpdateBook)
    api.Delete("/books/:id", handlers.RemoveBook)
	log.Fatal(app.Listen(":8080"))

 }

