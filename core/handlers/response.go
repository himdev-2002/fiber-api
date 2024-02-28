package handlers

import (
	"tde/fiber-api/core/structs"

	"github.com/gofiber/fiber/v2"
)

type ResponseParams struct {
	StatusCode int
	Message    string
	Data       any
}

type ResponsePagingParams struct {
	StatusCode int
	Message    string
	Paginate   *structs.Paginate
	Data       any
}

type NotFoundError struct {
	Message string
}

type BadRequestError struct {
	Message string
}

type InternalServerError struct {
	Message string
}

type UnauthorizedError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func (e *BadRequestError) Error() string {
	return e.Message
}

func (e *InternalServerError) Error() string {
	return e.Message
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}

func HandleError(ctx *fiber.Ctx, err error) error {
	var statusCode int

	switch err.(type) {
	case *NotFoundError:
		statusCode = fiber.StatusNotFound
	case *BadRequestError:
		statusCode = fiber.StatusBadRequest
	case *InternalServerError:
		statusCode = fiber.StatusInternalServerError
	case *UnauthorizedError:
		statusCode = fiber.StatusUnauthorized
	default:
		statusCode = fiber.StatusNotImplemented
	}

	return ResponseParams{
		StatusCode: statusCode,
		Message:    err.Error(),
	}.HandleResponse(ctx)
}

func (params ResponseParams) HandleResponse(ctx *fiber.Ctx) error {
	var response any
	status := false

	if params.StatusCode >= 200 && params.StatusCode < 300 {
		status = true
	}

	if params.Data != nil {
		response = &structs.ResponseWithoutPagingData{
			Code:    params.StatusCode,
			Status:  status,
			Message: params.Message,
			Data:    params.Data,
		}
	} else {
		response = &structs.ResponseWithoutData{
			Code:    params.StatusCode,
			Status:  status,
			Message: params.Message,
		}
	}

	return ctx.Status(params.StatusCode).JSON(response)
}

func (params ResponsePagingParams) HandleResponse(ctx *fiber.Ctx) error {
	var response any
	status := false

	if params.StatusCode >= 200 && params.StatusCode < 300 {
		status = true
	}

	if params.Data != nil {
		response = &structs.ResponseWithPagingData{
			Code:     params.StatusCode,
			Status:   status,
			Message:  params.Message,
			Paginate: params.Paginate,
			Data:     params.Data,
		}
	} else {
		response = &structs.ResponseWithoutData{
			Code:    params.StatusCode,
			Status:  status,
			Message: params.Message,
		}
	}

	return ctx.Status(params.StatusCode).JSON(response)
}
