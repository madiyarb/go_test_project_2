package main

import (
	_ "main/docs"

	"main/routes"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func main() {
	app := fiber.New()

	routes.Setup(app)

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Listen(":3000")
}
