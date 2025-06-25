package router

import (
	"him/fiber-api/api/sysparam/controllers"
	"him/fiber-api/core/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(r *fiber.App) {
	log, errLog := services.InfoLog()
	// sysparR := r.Group("/api/sysparam", middlewares.JWTAuthenticate())
	sysparR := r.Group("/api/sysparam")

	if errLog == nil {
		log.Info().Msgf("Add GET::/api/sysparam/user/:uid router")
	}
	sysparR.Get("/user/:uid", controllers.GetUserSysParams())

	if errLog == nil {
		log.Info().Msgf("Add POST::/api/sysparam router")
	}
	sysparR.Post("", controllers.AddSysParams())

	if errLog == nil {
		log.Info().Msgf("Add DELETE::/api/sysparam/:id router")
	}
	sysparR.Delete("/:id", controllers.RemoveSysParamByID())

	if errLog == nil {
		log.Info().Msgf("Add PUT::/api/sysparam/:id router")
	}
	sysparR.Put("/:id", controllers.UpdateSysParamByID())
}
