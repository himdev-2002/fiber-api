package models

import (
	"encoding/json"
	"him/fiber-api/core/services"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

const SysParamTbl string = "M_SYS_PARAM"

type SysParam struct {
	ID       uint64  `gorm:"primaryKey;autoIncrement" structs:"id" validated:"min=0"`
	CatID    uint64  `gorm:"not null;column:cat_id" structs:"cat_id" validate:"required,min=0;"`
	SubCatID uint64  `gorm:"column:_sub_cat_id" structs:"_sub_cat_id" validate:"min=0;"`
	Stat     uint8   `gorm:"default:1;precision:1;size:1;not null" structs:"stat" validate:"oneof=0 1"` // 0: inactive, 1: active
	IsSystem uint8   `gorm:"default:0;precision:1;size:1;not null;column:is_system" structs:"is_system" validate:"oneof=0 1"`
	Str1     string  `gorm:"size(255);column:_str_1" structs:"_str_1" validate:"max=255"`
	Str2     string  `gorm:"size(255);column:_str_2" structs:"_str_2" validate:"max=255"`
	Str3     string  `gorm:"size(255);column:_str_3" structs:"_str_3" validate:"max=255"`
	Num1     float64 `gorm:"column:_num_1" structs:"_num_1"`
	Num2     float64 `gorm:"column:_num_2" structs:"_num_2"`
	Num3     float64 `gorm:"column:_num_3" structs:"_num_3"`
	Bool1    uint8   `gorm:"precision:1;size:1;column:_bool_1" structs:"_bool_1" validate:"oneof=0 1"`
	Bool2    uint8   `gorm:"precision:1;size:1;column:_bool_2" structs:"_bool_2" validate:"oneof=0 1"`
	Bool3    uint8   `gorm:"precision:1;size:1;column:_bool_3" structs:"_bool_3" validate:"oneof=0 1"`

	CreatedDt time.Time `gorm:"not null;autoCreateTime;column:created_dt" structs:"created_dt"` // UNIX timestamp (int64)
	CreatedBy uint64    `gorm:"column:created_by" structs:"created_by"`

	UpdatedDt time.Time `gorm:"not null;autoUpdateTime;column:updated_dt" structs:"updated_dt"`
	UpdatedBy uint64    `gorm:"column:updated_by" structs:"updated_by"`

	LastSyncDt time.Time `gorm:"column:last_sync_dt" structs:"last_sync_dt"`
}

func (d SysParam) StatLabel() string {
	switch d.Stat {
	case 1:
		return "Active"
	case 0:
		return "Inactive"
	default:
		return "Unknown"
	}
}

func (d *SysParam) SaveSysParam() (*SysParam, error) {
	tx := services.DBCore.Session(&gorm.Session{SkipDefaultTransaction: true})
	err := tx.Table(SysParamTbl).Create(&d).Error
	if err != nil {
		return &SysParam{}, err
	}
	return d, nil
}

func GetSysParamByID(id *uint64) (*SysParam, error) {
	tx := services.DBCore.Session(&gorm.Session{PrepareStmt: true})

	var data SysParam
	if err := tx.Table(SysParamTbl).First(&data, id).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func GetSysParamByUID(uid *uint64, sysparams *[]map[string]interface{}) error {
	tx := services.DBCore.Session(&gorm.Session{PrepareStmt: true})
	err := tx.Table(SysParamTbl).Where("created_by = ?", *uid).
		Or("is_system = 1").Order("created_dt ASC").Find(sysparams).Error
	// fmt.Println(uid)
	// fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}

func RemoveSysParamByID(id *uint64) (*SysParam, error) {
	tx := services.DBCore.Session(&gorm.Session{SkipDefaultTransaction: true})

	// Cek dulu apakah user ada
	var data SysParam
	if err := tx.Table(SysParamTbl).First(&data, id).Error; err != nil {
		return nil, err
	}

	// Hapus user
	if err := tx.Table(SysParamTbl).Delete(&data).Error; err != nil {
		return &data, err
	}

	return &data, nil
}

func (d *SysParam) UpdateSysParam() (*SysParam, error) {
	tx := services.DBCore.Session(&gorm.Session{SkipDefaultTransaction: true})
	err := tx.Table(SysParamTbl).Save(&d).Error
	if err != nil {
		return &SysParam{}, err
	}
	return d, nil
}

// HOOK
func (d *SysParam) BeforeCreate(tx *gorm.DB) error {
	d.CreatedDt = time.Now()
	return nil
}

func (d *SysParam) BeforeSave(tx *gorm.DB) error {
	d.UpdatedDt = time.Now()
	d.Str1 = html.EscapeString(strings.TrimSpace(d.Str1))
	d.Str2 = html.EscapeString(strings.TrimSpace(d.Str2))
	d.Str3 = html.EscapeString(strings.TrimSpace(d.Str3))

	if d.Stat != 0 && d.Stat != 1 {
		d.Stat = 1 // default ke Active
	}
	return nil
}

func (d *SysParam) AfterCreate(tx *gorm.DB) error {
	nVal, err := json.Marshal(d)
	// fmt.Println("AFTER CREATE")
	// fmt.Println(string(nVal), err)
	if err != nil {
		return err
	}
	log := LogSysParam{
		PKID:      d.ID,
		FieldName: "_ALL_",
		CreatedBy: d.CreatedBy,
		CreatedDt: d.CreatedDt,
		OldVal:    "_ADD_",
		NewVal:    string(nVal),
	}
	log.SaveLogSysParam()
	return nil
}

func (d *SysParam) AfterDelete(tx *gorm.DB) (err error) {
	nVal, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}
	log := LogSysParam{
		PKID:      d.ID,
		FieldName: "_ALL_",
		CreatedBy: d.CreatedBy,
		CreatedDt: d.CreatedDt,
		OldVal:    string(nVal),
		NewVal:    "_REMOVE_",
	}
	log.SaveLogSysParam()
	return nil
}
