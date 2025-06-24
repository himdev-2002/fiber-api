package controllers

import (
	public_handlers "him/fiber-api/api/public/handlers"
	"him/fiber-api/api/public/structs"
	"him/fiber-api/core/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRegister() fiber.Handler {
	return func(c *fiber.Ctx) error {
		objReq := structs.UserRegisterRequest{}
		if err := c.BodyParser(&objReq); err == nil {

			if user, err := public_handlers.RegisterUser(&objReq); err == nil {
				resp := handlers.ResponseParams{
					StatusCode: fiber.StatusCreated,
					Message:    "created",
					Data:       user,
				}
				return resp.HandleResponse(c)
			} else {
				return handlers.HandleError(c, &handlers.BadRequestError{Message: err.Error()})
			}
		} else {
			return handlers.HandleError(c, &handlers.BadRequestError{Message: err.Error()})
		}
	}
}

func ChangePassword() fiber.Handler {
	return func(c *fiber.Ctx) error {
		objReq := structs.ChangePasswordRequest{}
		if err := c.BodyParser(&objReq); err == nil {

			if user, err := public_handlers.ChangePassword(&objReq); err == nil {
				resp := handlers.ResponseParams{
					StatusCode: fiber.StatusOK,
					Message:    "updated",
					Data:       user,
				}
				return resp.HandleResponse(c)
			} else {
				return handlers.HandleError(c, &handlers.BadRequestError{Message: err.Error()})
			}
		} else {
			return handlers.HandleError(c, &handlers.BadRequestError{Message: err.Error()})
		}
	}
}
