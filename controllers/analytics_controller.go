package controllers

import (
	"github.com/WeDias/golang-test-api/services"
	"github.com/gofiber/fiber/v2"
)

func BestSeller(ctx *fiber.Ctx) error {
	products, err := services.GetProducts()
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(products)
}
