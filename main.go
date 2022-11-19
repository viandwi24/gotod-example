package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/viandwi24/gotod-example/routes"
)

func main() {
	// init fiber app
	app := fiber.New()

	// load routes
	routes.InitRoutes(app)

	// start server
	log.Fatal(app.Listen(":3000"))
}
