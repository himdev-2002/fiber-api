package models

import (
	"him/fiber-api/core/services"
	"time"

	"gorm.io/gorm"
)

const LogCategoryTbl string = "LOG_SYS_PARAM_CAT"

type LogCategory struct {
	PKID      uint64    `gorm:"not null;index:lsysparcat_idx_1" structs:"pk_id" validated:"min=0;"`
	CreatedDt time.Time `gorm:"not null;autoCreateTime;column:created_dt" structs:"created_dt"` // UNIX timestamp (int64)
	CreatedBy uint64    `gorm:"column:created_by" structs:"created_by"`
	FieldName string    `gorm:"size:100;not null;column:fieldname" structs:"fieldname" validate:"required;max=100"`
	OldVal    string    `gorm:"type:text;column:_old_val" structs:"_old_val"`
	NewVal    string    `gorm:"type:text;column:new_val;not null" structs:"new_val" validate:"required"`

	LastSyncDt time.Time `gorm:"autoUpdateTime;column:last_sync_dt" structs:"last_sync_dt"`
}

func (l *LogCategory) SaveLogCategory() (*LogCategory, error) {
	tx := services.DBCore.Session(&gorm.Session{SkipDefaultTransaction: true})
	err := tx.Table(LogCategoryTbl).Create(&l).Error
	if err != nil {
		return &LogCategory{}, err
	}
	return l, nil
}

func (l *LogCategory) GetLatestByID(id *uint64) error {
	tx := services.DBCore.Session(&gorm.Session{PrepareStmt: true})
	err := tx.Table(LogCategoryTbl).Where("PKID = ?", *id).Order("updated_dt DESC").First(&l).Error
	if err != nil {
		return err
	}
	return nil
}
