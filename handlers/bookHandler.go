package handlers

import (
	"time"

	"github.com/CemHarput/bookStoreGo/config"
	"github.com/CemHarput/bookStoreGo/model"
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
    var books []model.Book

    config.Database.Find(&books)
    return c.Status(200).JSON(books)
}
func GetBook(c *fiber.Ctx) error {
    id := c.Params("id")
    var book model.Book

    result := config.Database.Find(&book, id)

    if result.RowsAffected == 0 {
        return c.SendStatus(404)
    }

    return c.Status(200).JSON(&book)
}
func AddBook(c *fiber.Ctx) error {
    book := new(model.Book)

    if err := c.BodyParser(book); err != nil {
        return c.Status(503).SendString(err.Error())
    }
    book.Model.CreatedAt=time.Now()

    config.Database.Create(&book)
    return c.Status(201).JSON(book)
}
func UpdateBook(c *fiber.Ctx) error {
    book := new(model.Book)
    id := c.Params("id")

    if err := c.BodyParser(book); err != nil {
        return c.Status(503).SendString(err.Error())
    }
    book.Model.UpdatedAt = time.Now()
    config.Database.Where("id = ?", id).Updates(&book)
    return c.Status(200).JSON(book)
}
func RemoveBook(c *fiber.Ctx) error {
    id := c.Params("id")
    var book model.Book

    result := config.Database.Delete(&book, id)

    if result.RowsAffected == 0 {
        return c.SendStatus(404)
    }
    return c.SendStatus(200)
}
