package handlers

import (
	"fmt"
	auth_structs "him/fiber-api/api/auth/structs"
	"him/fiber-api/core/models"
	"him/fiber-api/core/services"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func ValidateAccount(user *auth_structs.SignInRequest) (*models.User, error) {
	// Validate Param
	var validate = validator.New()
	validateErr := validate.Struct(user)
	if validateErr != nil {
		if log, err := services.DebugLog(); err == nil {
			log.Debug().Msgf(validateErr.Error())
		}
		return nil, fmt.Errorf("%v", validateErr.Error())
	}

	// if user.SecretKey != os.Getenv("JWT_SECRET_KEY") {
	// 	msg := "Invalid Key"
	// 	if log, err := services.DebugLog(); err == nil {
	// 		log.Debug().Msgf(msg)
	// 	}
	// 	return nil, fmt.Errorf("%v", msg)
	// }

	// Prepare Data
	u := models.User{}

	if err := u.FindByUsernameOrEmail(&user.Account); err != nil || u.ID == "" {
		msg := "User Not Found"
		if log, err := services.DebugLog(); err == nil {
			log.Debug().Msgf(msg)
		}
		return nil, fmt.Errorf("%v", msg)
	}

	// fmt.Println(u.Password)
	// fmt.Println(user.Password)
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	// fmt.Println(err)
	if err != nil {
		msg := "Invalid Authentication"
		if log, err := services.DebugLog(); err == nil {
			log.Debug().Msgf(msg)
		}
		return nil, fmt.Errorf("%v", msg)
	}

	return &u, nil
}
