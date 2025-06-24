package router

import (
	"him/fiber-api/api/auth/controllers"
	"him/fiber-api/core/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(r *fiber.App) {
	authR := r.Group("/api/auth")

	if log, err := services.InfoLog(); err == nil {
		log.Info().Msgf("Add POST::/api/auth/login router")
	}
	authR.Post("/login", controllers.SignIn())

	if log, err := services.InfoLog(); err == nil {
		log.Info().Msgf("Add POST::/api/auth/first-login router")
	}
	authR.Post("/first-login", controllers.FirstSignIn())

	// if log, err := services.InfoLog(); err == nil {
	// 	log.Info().Msgf("Add GET::/api/auth/logout router")
	// }
	// auth_r.Get("/logout", controllers.SignOut())
}
