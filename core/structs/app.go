package structs

import "github.com/go-playground/validator/v10"

type AppState struct {
	Validator *validator.Validate
	// tambahkan field lain seperti DB, logger, dsb.
}
