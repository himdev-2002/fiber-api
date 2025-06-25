package api

import (
	router "him/fiber-api/api/sysparam"
	"him/fiber-api/core/services"

	"github.com/gofiber/fiber/v2"
)

func SetupSysParamRouter(app *fiber.App) {
	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("Configuring 'sysparam' router...")
	}

	router.SetupRouter(app)

	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("'sysparam' router has been CONFIGURED!")
	}
}
