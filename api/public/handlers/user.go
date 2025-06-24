package handlers

import (
	"fmt"
	"os"
	"him/fiber-api/api/public/structs"
	"him/fiber-api/core/models"
	"him/fiber-api/core/services"

	"github.com/go-playground/validator/v10"
)

func RegisterUser(user *structs.UserRegisterRequest) (interface{}, error) {
	// Validate Param
	var validate = validator.New()
	validateErr := validate.Struct(user)
	if validateErr != nil {
		if log, err := services.DebugLog(); err == nil {
			log.Debug().Msgf(validateErr.Error())
		}
		return nil, fmt.Errorf("%v", validateErr.Error())
	}

	// Prepare Data
	u := models.User{
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  &user.LastName,
	}

	return u.SaveUser()
}

func ChangePassword(user *structs.ChangePasswordRequest) (interface{}, error) {
	// Validate Param
	var validate = validator.New()
	validateErr := validate.Struct(user)
	if validateErr != nil {
		if log, err := services.DebugLog(); err == nil {
			log.Debug().Msgf(validateErr.Error())
		}
		return nil, fmt.Errorf("%v", validateErr.Error())
	}

	if user.SecretKey != os.Getenv("SECRET_KEY") {
		msg := "Invalid Key"
		if log, err := services.DebugLog(); err == nil {
			log.Debug().Msgf(msg)
		}
		return nil, fmt.Errorf("%v", msg)
	}

	// Prepare Data
	u := models.User{}

	if err := u.FindByUsername(&user.Username); err != nil || u.ID == "" {
		msg := "User Not Found"
		if log, err := services.DebugLog(); err == nil {
			log.Debug().Msgf(msg)
		}
		return nil, fmt.Errorf("%v", msg)
	}

	return u.ChangePassword(user.Password)
}
