package router

import (
	"him/fiber-api/api/user/controllers"
	"him/fiber-api/core/helpers"
	"him/fiber-api/core/middlewares"
	"him/fiber-api/core/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(r *fiber.App) {
	usersR := r.Group("/api/users", middlewares.JWTAuthenticate())

	if log, err := services.InfoLog(); err == nil {
		log.Info().Msgf("Add GET::/api/users/active router")
	}
	usersR.Get("/active", controllers.GetActiveUsers())

	userR := r.Group("/api/user", middlewares.JWTAuthenticate())

	if log, err := services.InfoLog(); err == nil {
		log.Info().Msgf("Add GET::/api/user/:id router")
	}
	userR.Get("/:id", controllers.GetUserByID())

	if log, err := services.InfoLog(); err == nil {
		log.Info().Msgf("Add POST::/api/user router")
	}

	perms := helpers.NewSet[string]()
	perms.Add("ADMIN")
	userR.Post("", middlewares.RoutePermission(perms, helpers.IN), controllers.GetActiveUsers())
}
