package controllers

import (
	sysparam_handlers "him/fiber-api/api/sysparam/handlers"
	sysparam_structs "him/fiber-api/api/sysparam/structs"
	"him/fiber-api/core/handlers"
	"him/fiber-api/core/helpers"
	core_structs "him/fiber-api/core/structs"

	"github.com/gofiber/fiber/v2"
)

func GetUserSysParams() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// fmt.Println(c.Params("uid"))
		objReq := sysparam_structs.GetUserSysParamsRequest{
			UID: c.Params("uid"),
		}

		// state := c.Locals("state").(*core_structs.AppState)
		// if err := state.Validator.Struct(objReq); err == nil {
		if data, err := sysparam_handlers.GetUserSysParams(&objReq); err == nil {
			var resp *core_structs.DataResponse
			exclude := []string{}
			// exclude := []string{"password", "updated_at"}
			if resp, err = helpers.ConvertToDataResponse(data, &exclude); err == nil {
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
		// } else {
		// 	return handlers.HandleError(c, &handlers.BadRequestError{Message: err.Error()})
		// }
	}
}
