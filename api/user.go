package api

import (
	router "him/fiber-api/api/user"
	"him/fiber-api/core/services"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRouter(app *fiber.App) {
	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("Configuring 'user' router...")
	}

	router.SetupRouter(app)

	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("'user' router has been CONFIGURED!")
	}
}
