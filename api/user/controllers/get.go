package controllers

import (
	user_handlers "him/fiber-api/api/user/handlers"
	user_structs "him/fiber-api/api/user/structs"
	"him/fiber-api/core/handlers"
	"him/fiber-api/core/helpers"
	core_structs "him/fiber-api/core/structs"

	"github.com/fatih/structs"
	"github.com/gofiber/fiber/v2"
)

func GetActiveUsers() fiber.Handler {
	return func(c *fiber.Ctx) error {
		objReq := user_structs.GetActiveUsersRequest{}
		// fmt.Println(objReq)
		if err := c.BodyParser(&objReq); err == nil {
			if users, err := user_handlers.GetActiveUsers(&objReq); err == nil {
				var resp *core_structs.DataResponse
				exclude := []string{}
				// exclude := []string{"password", "updated_at"}
				if resp, err = helpers.ConvertToDataResponse(users, &exclude); err == nil {
					resp := handlers.ResponseParams{
						StatusCode: fiber.StatusOK,
						Message:    "ok",
						Data:       resp,
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

func GetUserByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		objReq := user_structs.GetUserByIDRequest{
			ID: c.Params("id"),
		}
		if user, err := user_handlers.GetUserByID(&objReq); err == nil {
			// user.Password = ""
			// user.UpdatedAt = time.Time{}
			m := structs.Map(user)
			resp := handlers.ResponseParams{
				StatusCode: fiber.StatusOK,
				Message:    "ok",
				Data:       m,
			}

			return resp.HandleResponse(c)
		} else {
			return handlers.HandleError(c, &handlers.BadRequestError{Message: err.Error()})
		}
	}
}
