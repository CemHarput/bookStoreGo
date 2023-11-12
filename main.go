package main

import (
	"log"

	"github.com/CemHarput/bookStoreGo/config"
	"github.com/gofiber/fiber/v2"
)
func main() {
	app := fiber.New()
	
	config.Connect()

	api := app.Group("/api/v1")

    api.Get("/books", handlers.getBooks)
    api.Get("/books/:id", handlers.getBook)
    api.Post("/books", handlers.addBook)
    api.Put("/books/:id", handlers.updateDog)
    api.Delete("/books/:id", handlers.removeDog)
	
	log.Fatal(app.Listen(":3000"))
 }