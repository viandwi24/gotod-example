package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viandwi24/gotod-example/app/models"
	"github.com/viandwi24/gotod-example/utils"
)

func GetProducts(c *fiber.Ctx) error {
	// get database
	products := []models.Product{
		{ID: 1, Title: "Product 1", Price: 1000},
		{ID: 2, Title: "Product 2", Price: 4000},
	}

	// return
	return utils.MakeResponse(c, products, fiber.StatusOK)
}
