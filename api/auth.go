package api

import (
	router "tde/fiber-api/api/auth"
	"tde/fiber-api/core/services"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRouter(app *fiber.App) {
	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("Configuring 'auth' router...")
	}

	router.SetupRouter(app)

	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("'auth' router has been CONFIGURED!")
	}
}
