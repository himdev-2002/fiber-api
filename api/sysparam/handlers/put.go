package handlers

import (
	"fmt"
	"him/fiber-api/api/sysparam/structs"
	"him/fiber-api/core/models"
	"him/fiber-api/core/services"

	"github.com/jinzhu/copier"
)

func UpdateSysParamsByID(param *structs.PutUserSysParamsRequest) (*models.SysParam, error) {
	// Prepare Data
	var id uint64
	fmt.Sscanf(param.ID, "%d", &id)
	// fmt.Println(id)
	ret, err := models.GetSysParamByID(&id)
	fmt.Println(ret, err)
	if err != nil {
		msg := err.Error()
		if log, err := services.DebugLog(); err == nil {
			log.Debug().Msgf(msg)
		}
		return nil, fmt.Errorf("%v", msg)
	}
	// fmt.Println(ret)

	copier.Copy(&ret, &param)
	ret2, err := ret.UpdateSysParam()
	if err != nil {
		msg := err.Error()
		if log, err := services.DebugLog(); err == nil {
			log.Debug().Msgf(msg)
		}
		return nil, fmt.Errorf("%v", msg)
	}
	return ret2, nil
}
