package controllers

import (
	sysparam_handlers "him/fiber-api/api/sysparam/handlers"
	sysparam_structs "him/fiber-api/api/sysparam/structs"
	"him/fiber-api/core/handlers"

	"github.com/fatih/structs"
	"github.com/gofiber/fiber/v2"
)

func RemoveSysParamByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// fmt.Println(c.Params("uid"))
		objReq := sysparam_structs.DeleteSysParamsRequest{
			ID: c.Params("id"),
		}

		// state := c.Locals("state").(*core_structs.AppState)
		// if err := state.Validator.Struct(objReq); err == nil {
		if data, err := sysparam_handlers.DeleteSysParamsByID(&objReq); err == nil {
			d := structs.Map(data)
			resp := handlers.ResponseParams{
				StatusCode: fiber.StatusOK,
				Message:    "ok",
				Data:       d,
			}
			return resp.HandleResponse(c)
		} else {
			return handlers.HandleError(c, &handlers.BadRequestError{Message: err.Error()})
		}
		// } else {
		// 	return handlers.HandleError(c, &handlers.BadRequestError{Message: err.Error()})
		// }
	}
}
