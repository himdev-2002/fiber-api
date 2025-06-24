package controllers

import (
	"fmt"
	auth_handlers "him/fiber-api/api/auth/handlers"
	auth_structs "him/fiber-api/api/auth/structs"
	"him/fiber-api/core/handlers"
	"him/fiber-api/core/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid/v5"
	"github.com/lithammer/shortuuid"
)

func SignIn() fiber.Handler {
	return func(c *fiber.Ctx) error {
		objReq := auth_structs.SignInRequest{}
		if err := c.BodyParser(&objReq); err == nil {
			if user, err := auth_handlers.ValidateAccount(&objReq); err == nil {
				if token, refresh, errToken := helpers.GenerateToken(user); errToken == nil {
					ret := map[string]interface{}{
						"t":  token,
						"rt": refresh,
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

func FirstSignIn() fiber.Handler {
	return func(c *fiber.Ctx) error {
		objReq := auth_structs.SignInRequest{}
		if err := c.BodyParser(&objReq); err == nil {
			if user, err := auth_handlers.ValidateAccount(&objReq); err == nil {
				uuid_v7, err := uuid.NewV7()
				if err != nil {
					return handlers.HandleError(c, &handlers.InternalServerError{Message: err.Error()})
				}
				s_uuid := shortuuid.New()
				if _, err := user.ChangeToken(fmt.Sprintf("%s_%s!", s_uuid, uuid_v7)); err == nil {
					resp := handlers.ResponseParams{
						StatusCode: fiber.StatusOK,
						Message:    "ok",
						Data:       fmt.Sprintf("%s_%s!", s_uuid, uuid_v7),
					}
					return resp.HandleResponse(c)
				} else {
					return handlers.HandleError(c, &handlers.InternalServerError{Message: err.Error()})
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
