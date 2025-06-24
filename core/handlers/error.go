package handlers

import (
	"errors"
	"him/fiber-api/core/services"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	msg := err.Error()
	ret := map[string]any{
		"s": false,
		"m": msg,
	}
	err = ctx.Status(code).JSON(ret)
	if err != nil {
		if log, err := services.ErrorLog(); err == nil {
			log.Log().Msgf(msg)
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(ret)
	}

	// Return from handler
	return nil
}
