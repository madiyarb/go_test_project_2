package handlers

import (
	"main/entities"

	"github.com/gofiber/fiber/v2"
)

// GetProfile godoc
// @Summary Get profile
// @Description Returns Hello World
// @Tags profile
// @Produce plain
// @Success 200 {string} string
// @Router /profile [get]
func GetProfile(c *fiber.Ctx) error {
	return c.SendString("Hello from controller")
}

// @Summary Create Profile
// @Description returns created profile
// @Tags profile
// @Produce json
// @Success 200 {string} UserDTO
// @Router /profile [post]
func CreateProfile(c *fiber.Ctx) error {
	userDTO := entities.NewUserDTO(1, "Madi", "email@example.com")
	return c.JSON(userDTO)
}
