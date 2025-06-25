package controllers

import (
	sysparam_handlers "him/fiber-api/api/sysparam/handlers"
	sysparam_structs "him/fiber-api/api/sysparam/structs"
	"him/fiber-api/core/handlers"

	"github.com/fatih/structs"
	"github.com/gofiber/fiber/v2"
)

func UpdateSysParamByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var objReq sysparam_structs.PutUserSysParamsRequest
		if err := c.BodyParser(&objReq); err == nil {
			objReq.ID = c.Params("id")
			if data, err := sysparam_handlers.UpdateSysParamsByID(&objReq); err == nil {
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
		} else {
			return handlers.HandleError(c, &handlers.BadRequestError{Message: err.Error()})
		}
		// objReq := sysparam_structs.PutUserSysParamsRequest{
		// 	ID: c.Params("id"),
		// }
		// fmt.Println(structs.Map(objReq))

		// state := c.Locals("state").(*core_structs.AppState)
		// if err := state.Validator.Struct(objReq); err == nil {
		// fmt.Println(structs.Map(objReq))
		// if data, err := sysparam_handlers.UpdateSysParamsByID(&objReq); err == nil {
		// 	d := structs.Map(data)
		// 	resp := handlers.ResponseParams{
		// 		StatusCode: fiber.StatusOK,
		// 		Message:    "ok",
		// 		Data:       d,
		// 	}
		// 	return resp.HandleResponse(c)
		// } else {
		// 	return handlers.HandleError(c, &handlers.BadRequestError{Message: err.Error()})
		// }
		// } else {
		// 	// fmt.Println(structs.Map(objReq))
		// 	return handlers.HandleError(c, &handlers.BadRequestError{Message: err.Error()})
		// }
	}
}
