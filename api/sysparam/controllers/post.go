package controllers

import (
	"fmt"
	sysparam_handlers "him/fiber-api/api/sysparam/handlers"
	sysparam_structs "him/fiber-api/api/sysparam/structs"
	"him/fiber-api/core/handlers"

	"github.com/fatih/structs"
	"github.com/gofiber/fiber/v2"
)

func AddSysParams() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// fmt.Println(c.Params("uid"))
		var objReq sysparam_structs.PostUserSysParamsRequest
		// fmt.Println(objReq)
		if err := c.BodyParser(&objReq); err == nil {
			// state := c.Locals("state").(*core_structs.AppState)
			// fmt.Println(structs.Map(objReq))
			// state.Validator = validator.New()

			// err := state.Validator.Struct(objReq)
			// if err := state.Validator.Struct(objReq); err == nil {
			// 	fmt.Println("AddUserSysParams")
			if data, err := sysparam_handlers.AddUserSysParams(&objReq); err == nil {
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
			fmt.Println(err)
			return handlers.HandleError(c, &handlers.BadRequestError{Message: err.Error()})
			// }
		} else {
			return handlers.HandleError(c, &handlers.BadRequestError{Message: err.Error()})
		}
	}
}
