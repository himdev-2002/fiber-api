package core

import (
	"log"
	"tde/fiber-api/core/setup"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func SetupApp() *fiber.App {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	setup.SetupDBSys()
	fiberEngine := setup.SetupFiber()
	setup.SetupRouter(fiberEngine)

	return fiberEngine
}
