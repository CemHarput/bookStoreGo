package handlers

import (
	"time"

	"github.com/CemHarput/bookStoreGo/config"
	"github.com/CemHarput/bookStoreGo/model"
	"github.com/gofiber/fiber/v2"
)
func GetPurchaseOrders(c *fiber.Ctx) error {
    var orders []model.PurchaseOrder

    config.Database.Find(&orders)
    return c.Status(200).JSON(orders)
}
func GetPurchaseOrder(c *fiber.Ctx) error {
    id := c.Params("id")
    var purchaseOrder model.PurchaseOrder

    result := config.Database.Find(&purchaseOrder, id)

    if result.RowsAffected == 0 {
        return c.SendStatus(404)
    }

    return c.Status(200).JSON(&purchaseOrder)
}
func AddPurchaseOrder(c *fiber.Ctx) error {
    purchaseOrder := new(model.PurchaseOrder)

    if err := c.BodyParser(purchaseOrder); err != nil {
        return c.Status(503).SendString(err.Error())
    }
    purchaseOrder.Model.CreatedAt=time.Now()

    config.Database.Create(&purchaseOrder)
    return c.Status(201).JSON(purchaseOrder)
}
func UpdatePurchaseOrder(c *fiber.Ctx) error {
    purchaseOrder := new(model.PurchaseOrder)
    id := c.Params("id")

    if err := c.BodyParser(purchaseOrder); err != nil {
        return c.Status(503).SendString(err.Error())
    }
    purchaseOrder.Model.UpdatedAt = time.Now()
    config.Database.Where("id = ?", id).Updates(&purchaseOrder)
    return c.Status(200).JSON(purchaseOrder)
}