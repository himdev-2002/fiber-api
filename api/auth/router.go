package router

import (
	"tde/fiber-api/api/auth/controllers"
	"tde/fiber-api/core/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(r *fiber.App) {
	auth_r := r.Group("/auth")

	if log, err := services.InfoLog(); err == nil {
		log.Info().Msgf("Add POST::/auth/login router")
	}
	auth_r.Post("/login", controllers.SignIn())

	// if log, err := services.InfoLog(); err == nil {
	// 	log.Info().Msgf("Add GET::/auth/logout router")
	// }
	// auth_r.Get("/logout", controllers.SignOut())
}
