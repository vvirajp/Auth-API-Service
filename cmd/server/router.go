package server

import (
	"API/components/auth"

	"github.com/gofiber/fiber/v2"
)

func LoadRoutes(app *fiber.App) {
	auth.Routes(app)
}
