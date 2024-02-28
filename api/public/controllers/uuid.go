package controllers

import (
	"tde/fiber-api/core/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/lithammer/shortuuid"
)

func ShortUUID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		u := shortuuid.New()
		resp := handlers.ResponseParams{
			StatusCode: fiber.StatusOK,
			Message:    "ok",
			Data:       u,
		}

		return resp.HandleResponse(c)
	}
}
