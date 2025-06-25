package handlers

import (
	"fmt"
	"him/fiber-api/api/sysparam/structs"
	"him/fiber-api/core/models"
	"him/fiber-api/core/services"
)

func GetUserSysParams(param *structs.GetUserSysParamsRequest) (*[]map[string]interface{}, error) {
	// Prepare Data
	var data []map[string]interface{}
	var uid uint64
	fmt.Sscanf(param.UID, "%d", &uid)
	// fmt.Println(uid)
	if err := models.GetSysParamByUID(&uid, &data); err != nil {
		msg := "User System Params Not Found"
		if log, err := services.DebugLog(); err == nil {
			log.Debug().Msgf(msg)
		}
		return nil, fmt.Errorf("%v", msg)
	}

	return &data, nil
}
