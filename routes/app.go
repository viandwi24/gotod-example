package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/viandwi24/gotod-example/app/controllers"
)

func InitRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// products
	route.Get("/products", controllers.GetProducts)
}
