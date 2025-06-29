package api

import (
	router "him/fiber-api/api/public"
	"him/fiber-api/core/services"

	"github.com/gofiber/fiber/v2"
)

func SetupPublicRouter(app *fiber.App) {
	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("Configuring 'public' router...")
	}

	if log, err := services.InfoLog(); err == nil {
		log.Info().Msgf("Add GET::/ping router")
	}
	app.Get("/api/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("pong")
	})

	router.SetupRouter(app)

	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("'public' router has been CONFIGURED!")
	}
}
