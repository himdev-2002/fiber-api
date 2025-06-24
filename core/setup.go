package core

import (
	"embed"
	"him/fiber-api/core/setup"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func SetupApp(embedFile embed.FS, dbPath string) *fiber.App {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	setup.SetupDBSys(dbPath)
	fiberEngine := setup.SetupFiber(embedFile)
	setup.SetupRouter(fiberEngine)

	return fiberEngine
}
