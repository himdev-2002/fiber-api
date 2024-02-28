package router

import (
	"tde/fiber-api/api/public/controllers"
	"tde/fiber-api/core/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	g_r := app.Group("/generate")

	if log, err := services.InfoLog(); err == nil {
		log.Info().Msgf("Add GET::/generate/short-uuid router")
	}
	g_r.Get("/short-uuid", controllers.ShortUUID())

	if log, err := services.InfoLog(); err == nil {
		log.Info().Msgf("Add POST::/generate/bcrypt router")
	}
	g_r.Post("/bcrypt", controllers.Bcrypt())

	user_r := app.Group("/user")

	if log, err := services.InfoLog(); err == nil {
		log.Info().Msgf("Add POST::/user/register router")
	}
	user_r.Post("/register", controllers.UserRegister())

	if log, err := services.InfoLog(); err == nil {
		log.Info().Msgf("Add PUT::/user/change-password router")
	}
	user_r.Put("/change-password", controllers.ChangePassword())
}
