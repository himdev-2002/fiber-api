package router

import (
	"tde/fiber-api/api/user/controllers"
	"tde/fiber-api/core/middlewares"
	"tde/fiber-api/core/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(r *fiber.App) {
	users_r := r.Group("/users")

	if log, err := services.InfoLog(); err == nil {
		log.Info().Msgf("Add GET::/users/active router")
	}
	users_r.Get("/active", middlewares.JWTAuthenticate(), controllers.GetActiveUsers())

	user_r := r.Group("/user")

	if log, err := services.InfoLog(); err == nil {
		log.Info().Msgf("Add GET::/user/:id router")
	}
	user_r.Get("/:id", middlewares.JWTAuthenticate(), controllers.GetUserByID())
}
