package controllers

import (
	"him/fiber-api/core/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid/v5"
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

func UUID_V7() fiber.Handler {
	return func(c *fiber.Ctx) error {
		u, err := uuid.NewV7()
		if err != nil {
			return handlers.HandleError(c, &handlers.InternalServerError{Message: err.Error()})
		}
		resp := handlers.ResponseParams{
			StatusCode: fiber.StatusOK,
			Message:    "ok",
			Data:       u.String(),
		}

		return resp.HandleResponse(c)
	}
}
