package handlers

import (
	"fmt"
	"him/fiber-api/api/sysparam/structs"
	"him/fiber-api/core/models"
	"him/fiber-api/core/services"

	"github.com/jinzhu/copier"
)

func AddUserSysParams(param *structs.PostUserSysParamsRequest) (*models.SysParam, error) {
	// Prepare Data
	d := models.SysParam{}
	copier.Copy(&d, &param)
	d.CreatedBy = 999
	d.UpdatedBy = 999

	// fmt.Println(d)
	ret, err := d.SaveSysParam()
	if err != nil {
		msg := err.Error()
		if log, err := services.DebugLog(); err == nil {
			log.Debug().Msgf(msg)
		}
		return nil, fmt.Errorf("%v", msg)
	}

	return ret, nil
}
