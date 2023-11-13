package main

import (
	"log"

	"github.com/CemHarput/bookStoreGo/config"
	"github.com/CemHarput/bookStoreGo/docs"
	"github.com/CemHarput/bookStoreGo/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)
func main() {
	app := fiber.New()
	
	config.Connect()

	api := app.Group("/api/v1")

    api.Get("/books", handlers.GetBooks)
    api.Get("/books/:id", handlers.GetBook)
    api.Post("/books", handlers.AddBook)
    api.Put("/books/:id", handlers.UpdateBook)
    api.Delete("/books/:id", handlers.RemoveBook)
	docs.SwaggerInfo.ReadDoc()
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))
	log.Fatal(app.Listen(":3000"))
 }

