package controllers

import (
	auth_handlers "tde/fiber-api/api/auth/handlers"
	auth_structs "tde/fiber-api/api/auth/structs"
	"tde/fiber-api/core/handlers"
	"tde/fiber-api/core/helpers"

	"github.com/gofiber/fiber/v2"
)

func SignIn() fiber.Handler {
	return func(c *fiber.Ctx) error {
		objReq := auth_structs.SignInRequest{}
		if err := c.BodyParser(&objReq); err == nil {
			if user, err := auth_handlers.ValidateAccount(&objReq); err == nil {
				if token, refresh, errToken := helpers.GenerateToken(user); errToken == nil {
					ret := map[string]interface{}{
						"token":         token,
						"refresh_token": refresh,
					}
					resp := handlers.ResponseParams{
						StatusCode: fiber.StatusOK,
						Message:    "ok",
						Data:       ret,
					}
					return resp.HandleResponse(c)
				} else {
					return handlers.HandleError(c, &handlers.InternalServerError{Message: errToken.Error()})
				}
			} else {
				return handlers.HandleError(c, &handlers.BadRequestError{Message: err.Error()})
			}
		} else {
			return handlers.HandleError(c, &handlers.BadRequestError{Message: err.Error()})
		}
	}
}

// func SignOut() fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		objReq := user_structs.GetUserByIDRequest{
// 			ID: c.Params("id"),
// 		}
// 		if user, err := user_handlers.GetUserByID(&objReq); err == nil {
// 			user.Password = ""
// 			user.UpdatedAt = time.Time{}
// 			m := structs.Map(user)
// 			resp := handlers.ResponseParams{
// 				StatusCode: fiber.StatusOK,
// 				Message:    "ok",
// 				Data:       m,
// 			}

// 			return resp.HandleResponse(c)
// 		} else {
// 			return handlers.HandleError(c, &handlers.BadRequestError{Message: err.Error()})
// 		}
// 	}
// }

// func RefreshToken() fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		objReq := user_structs.GetUserByIDRequest{
// 			ID: c.Params("id"),
// 		}
// 		if user, err := user_handlers.GetUserByID(&objReq); err == nil {
// 			user.Password = ""
// 			user.UpdatedAt = time.Time{}
// 			m := structs.Map(user)
// 			resp := handlers.ResponseParams{
// 				StatusCode: fiber.StatusOK,
// 				Message:    "ok",
// 				Data:       m,
// 			}

// 			return resp.HandleResponse(c)
// 		} else {
// 			return handlers.HandleError(c, &handlers.BadRequestError{Message: err.Error()})
// 		}
// 	}
// }
