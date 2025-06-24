package handlers

import (
	"fmt"
	"him/fiber-api/api/user/structs"
	"him/fiber-api/core/models"
	"him/fiber-api/core/services"

	"github.com/go-playground/validator/v10"
)

func GetActiveUsers(param *structs.GetActiveUsersRequest) (*[]map[string]interface{}, error) {
	// Validate Param
	var validate = validator.New()
	validateErr := validate.Struct(param)
	if validateErr != nil {
		if log, err := services.DebugLog(); err == nil {
			log.Debug().Msgf(validateErr.Error())
		}
		return nil, fmt.Errorf("%v", validateErr.Error())
	}

	// Prepare Data
	var users []map[string]interface{}
	var like_cond []string
	like_cond = append(like_cond, "status=@status")

	like_value := make(map[string]interface{})
	like_value["status"] = 1

	if param.EmailLike != "" {
		like_cond = append(like_cond, "email like @email")
		like_value["email"] = fmt.Sprintf("%%%v%%", param.EmailLike)
	}
	// fmt.Println(like_cond, like_value)

	if err := models.SearchActiveUsers(&like_cond, &like_value, &users); err != nil {
		msg := "Failed retrieve data"
		if log, err := services.DebugLog(); err == nil {
			log.Debug().Msgf(msg)
		}
		return nil, fmt.Errorf("%v", msg)
	}

	return &users, nil
}

func GetUserByID(param *structs.GetUserByIDRequest) (*models.User, error) {
	// Validate Param
	var validate = validator.New()
	validateErr := validate.Struct(param)
	if validateErr != nil {
		if log, err := services.DebugLog(); err == nil {
			log.Debug().Msgf(validateErr.Error())
		}
		return nil, fmt.Errorf("%v", validateErr.Error())
	}

	// Prepare Data
	u := models.User{}

	if err := u.FindByID(&param.ID); err != nil || u.ID == "" {
		msg := "User Not Found"
		if log, err := services.DebugLog(); err == nil {
			log.Debug().Msgf(msg)
		}
		return nil, fmt.Errorf("%v", msg)
	}

	return &u, nil
}
