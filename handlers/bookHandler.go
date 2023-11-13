package handlers

import (
	"time"

	"github.com/CemHarput/bookStoreGo/config"
	"github.com/CemHarput/bookStoreGo/model"
	"github.com/gofiber/fiber/v2"
)

// GetBooks godoc
// @Summary Get all books
// @Description Get all books available
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {array} Book
// @Router /books [get]
func GetBooks(c *fiber.Ctx) error {
    var books []model.Book

    config.Database.Find(&books)
    return c.Status(200).JSON(books)
}
// GetBook godoc
// @Summary Get a book by ID
// @Description Get a book by its ID
// @Tags Books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} Book
// @Router /books/{id} [get]
func GetBook(c *fiber.Ctx) error {
    id := c.Params("id")
    var book model.Book

    result := config.Database.Find(&book, id)

    if result.RowsAffected == 0 {
        return c.SendStatus(404)
    }

    return c.Status(200).JSON(&book)
}
// AddBook godoc
// @Summary Add a new book
// @Description Add a new book to the store
// @Tags Books
// @Accept json
// @Produce json
// @Param data body Book true "Book object"
// @Success 200 {object} Book
// @Router /books [post]
func AddBook(c *fiber.Ctx) error {
    book := new(model.Book)

    if err := c.BodyParser(book); err != nil {
        return c.Status(503).SendString(err.Error())
    }
    book.Model.CreatedAt=time.Now()

    config.Database.Create(&book)
    return c.Status(201).JSON(book)
}
// UpdateBook godoc
// @Summary Update a book by ID
// @Description Update a book's details by its ID
// @Tags Books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param data body Book true "Book object"
// @Success 200 {object} Book
// @Router /books/{id} [put]
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
// RemoveBook godoc
// @Summary Remove a book by ID
// @Description Remove a book by its ID
// @Tags Books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 204
// @Router /books/{id} [delete]
func RemoveBook(c *fiber.Ctx) error {
    id := c.Params("id")
    var book model.Book

    result := config.Database.Delete(&book, id)

    if result.RowsAffected == 0 {
        return c.SendStatus(404)
    }
    return c.SendStatus(200)
}
