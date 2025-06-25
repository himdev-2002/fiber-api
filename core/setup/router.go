package setup

import (
	"him/fiber-api/api"
	"him/fiber-api/core/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(fiberEngine *fiber.App) {
	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("Configuring app router...")
	}

	api.SetupPublicRouter(fiberEngine)
	// api.SetupAuthRouter(fiberEngine)
	// api.SetupUserRouter(fiberEngine)
	api.SetupSysParamRouter(fiberEngine)
	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("App router has been CONFIGURED!")
	}
}
