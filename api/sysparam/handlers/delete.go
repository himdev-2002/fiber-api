package handlers

import (
	"fmt"
	"him/fiber-api/api/sysparam/structs"
	"him/fiber-api/core/models"
	"him/fiber-api/core/services"
)

func DeleteSysParamsByID(param *structs.DeleteSysParamsRequest) (*models.SysParam, error) {
	// Prepare Data
	var id uint64
	fmt.Sscanf(param.ID, "%d", &id)
	// fmt.Println(id)
	ret, err := models.RemoveSysParamByID(&id)
	if err != nil {
		msg := err.Error()
		if log, err := services.DebugLog(); err == nil {
			log.Debug().Msgf(msg)
		}
		return nil, fmt.Errorf("%v", msg)
	}

	return ret, nil
}
