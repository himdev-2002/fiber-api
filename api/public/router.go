package router

import (
	"him/fiber-api/api/public/controllers"
	"him/fiber-api/core/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	gR := app.Group("/api/generate")

	if log, err := services.InfoLog(); err == nil {
		log.Info().Msgf("Add GET::/api/generate/short-uuid router")
	}
	gR.Get("/short-uuid", controllers.ShortUUID())

	if log, err := services.InfoLog(); err == nil {
		log.Info().Msgf("Add GET::/api/generate/uuid-v7 router")
	}
	gR.Get("/uuid-v7", controllers.UUID_V7())

	if log, err := services.InfoLog(); err == nil {
		log.Info().Msgf("Add POST::/api/generate/bcrypt router")
	}
	gR.Post("/bcrypt", controllers.Bcrypt())

	userR := app.Group("/api/user")

	if log, err := services.InfoLog(); err == nil {
		log.Info().Msgf("Add POST::/api/user/register router")
	}
	userR.Post("/register", controllers.UserRegister())

	if log, err := services.InfoLog(); err == nil {
		log.Info().Msgf("Add PUT::/api/user/change-password router")
	}
	userR.Put("/change-password", controllers.ChangePassword())
}
