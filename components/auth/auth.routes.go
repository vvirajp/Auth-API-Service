package auth

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) fiber.Router {
	router := app.Group("/api/assignment")
	router.Post("/auth/signup", SignupController)
	router.Post("/auth/signin", SigninController)
	router.Post("/auth/verify", VerifyController)
	router.Post("/auth/refreshToken", RefreshController)
	router.Post("/auth/revokeToken", RevokeTokenController)
	return router
}
