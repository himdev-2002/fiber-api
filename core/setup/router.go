package setup

import (
	"tde/fiber-api/api"
	"tde/fiber-api/core/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(fiberEngine *fiber.App) {
	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("Configuring app router...")
	}

	api.SetupPublicRouter(fiberEngine)
	api.SetupAuthRouter(fiberEngine)
	api.SetupUserRouter(fiberEngine)
	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("App router has been CONFIGURED!")
	}
}
