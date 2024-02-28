package controllers

import (
	"tde/fiber-api/api/public/structs"
	"tde/fiber-api/core/handlers"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Bcrypt() fiber.Handler {
	return func(c *fiber.Ctx) error {
		objReq := structs.GenerateBcryptRequest{}
		if err := c.BodyParser(&objReq); err == nil {
			if hash, err := bcrypt.GenerateFromPassword([]byte(objReq.Password), bcrypt.DefaultCost); err == nil {
				ret := map[string]interface{}{
					"pass": objReq.Password,
					"hash": string(hash),
				}
				resp := handlers.ResponseParams{
					StatusCode: fiber.StatusOK,
					Message:    "ok",
					Data:       ret,
				}
				return resp.HandleResponse(c)
			} else {
				return handlers.HandleError(c, &handlers.InternalServerError{Message: err.Error()})
			}
		} else {
			return handlers.HandleError(c, &handlers.BadRequestError{Message: err.Error()})
		}
	}
}
