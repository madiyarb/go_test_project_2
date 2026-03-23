package routes

import (
	"main/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/profile", handlers.GetProfile)
	app.Post("/profile", handlers.CreateProfile)
}
