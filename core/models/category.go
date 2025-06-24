package models

import (
	"encoding/json"
	"fmt"
	"him/fiber-api/core/services"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

const CategoryTbl string = "M_SYS_PARAM_CAT"

type Category struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement" structs:"id" validated:"min=0;"`
	Name     string `gorm:"size:100;not null" structs:"name" validate:"required;max=100"`
	Stat     uint8  `gorm:"default:1;precision:1;size:1;not null" structs:"stat" validate:"oneof=0 1"`      // 0: inactive, 1: active
	IsSystem uint8  `gorm:"default:1;precision:1;size:1;not null" structs:"is_system" validate:"oneof=0 1"` // 0: inactive, 1: active

	CreatedDt time.Time `gorm:"not null;autoCreateTime;column:created_dt" structs:"created_dt"` // UNIX timestamp (int64)
	CreatedBy uint64    `gorm:"column:created_by" structs:"created_by"`

	UpdatedDt time.Time `gorm:"not null;autoUpdateTime;column:updated_dt" structs:"updated_dt"`
	UpdatedBy uint64    `gorm:"column:updated_by" structs:"updated_by"`

	LastSyncDt time.Time `gorm:"autoUpdateTime;column:last_sync_dt" structs:"last_sync_dt"`
}

func (c Category) StatLabel() string {
	switch c.Stat {
	case 1:
		return "Active"
	case 0:
		return "Inactive"
	default:
		return "Unknown"
	}
}

func (c *Category) SaveCategory() (*Category, error) {
	tx := services.DBCore.Session(&gorm.Session{SkipDefaultTransaction: true})
	err := tx.Table(CategoryTbl).Create(&c).Error
	if err != nil {
		return &Category{}, err
	}
	return c, nil
}

// HOOK
func (c *Category) BeforeCreate(tx *gorm.DB) error {
	c.Name = html.EscapeString(strings.TrimSpace(c.Name))
	c.CreatedDt = time.Now()

	if c.Stat != 0 && c.Stat != 1 {
		c.Stat = 1 // default ke Active
	}
	return nil
}

func (c *Category) BeforeSave(tx *gorm.DB) error {
	c.Name = html.EscapeString(strings.TrimSpace(c.Name))
	c.UpdatedDt = time.Now()

	if c.Stat != 0 && c.Stat != 1 {
		c.Stat = 1 // default ke Active
	}
	return nil
}

func (c *Category) AfterCreate(tx *gorm.DB) error {
	// fmt.Println("[HOOK] AfterCreate - category created with ID:", c.ID)
	nVal, err := json.MarshalIndent(c, "", "  ")
	if err == nil {
		return err
	}
	log := LogCategory{
		PKID:      c.ID,
		FieldName: "_ALL_",
		CreatedBy: c.CreatedBy,
		CreatedDt: c.CreatedDt,
		OldVal:    "_ADD_",
		NewVal:    string(nVal),
	}
	log.SaveLogCategory()
	return nil
}

func (c *Category) AfterDelete(tx *gorm.DB) (err error) {
	// fmt.Println("[HOOK] AfterDelete - deleted category ID:", c.ID)
	nVal, err := json.MarshalIndent(c, "", "  ")
	if err == nil {
		return err
	}
	log := LogCategory{
		PKID:      c.ID,
		FieldName: "_ALL_",
		CreatedBy: c.CreatedBy,
		CreatedDt: c.CreatedDt,
		OldVal:    string(nVal),
		NewVal:    "_REMOVE_",
	}
	log.SaveLogCategory()
	return nil
}

func (c *Category) addUpdateLog(old *Category) error {
	// Bandingkan field tertentu
	changes := compareFields(&old, c, []string{},
		[]string{"ID", "CreatedDt", "CreatedBy", "UpdatedDt", "UpdatedBy", "LastSyncDt"})

	// Simpan perubahan ke log DB
	for field, val := range changes {
		log := LogCategory{
			PKID:      c.ID,
			FieldName: field,
			CreatedBy: c.CreatedBy,
			CreatedDt: c.CreatedDt,
			OldVal:    fmt.Sprintf("%v", val[0]),
			NewVal:    fmt.Sprintf("%v", val[1]),
		}
		log.SaveLogCategory()
	}

	return nil
}

// func (c *Category) BeforeUpdate(tx *gorm.DB) error {
// 	c.Name = html.EscapeString(strings.TrimSpace(c.Name))
// 	c.UpdatedDt = time.Now()

// 	if c.Stat != 0 && c.Stat != 1 {
// 		c.Stat = 1 // default ke Active
// 	}

// 	// Ambil data lama
// 	var old Category
// 	if err := tx.Unscoped().First(&old, c.ID).Error; err != nil {
// 		return nil // skip silently
// 	}

// 	// Bandingkan field tertentu
// 	changes := compareFields(&old, c, []string{},
// 		[]string{"ID", "CreatedDt", "CreatedBy", "UpdatedDt", "UpdatedBy", "LastSyncDt"})

// 	// Simpan perubahan ke log DB
// 	for field, val := range changes {
// 		log := LogCategory{
// 			PKID:      c.ID,
// 			FieldName: field,
// 			CreatedBy: c.CreatedBy,
// 			CreatedDt: c.CreatedDt,
// 			OldVal:    fmt.Sprintf("%v", val[0]),
// 			NewVal:    fmt.Sprintf("%v", val[1]),
// 		}
// 		log.SaveLogCategory()
// 	}

// 	return nil
// }

// func (c *Category) AfterUpdate(tx *gorm.DB) error {
// 	log := LogCategory{}
// 	if err := log.GetLatestByPKID(&c.ID); err == nil {

// 	}

// 	// var cat Category
// 	// db.Order("updated_dt DESC").First(&cat)

// 	c.Name = html.EscapeString(strings.TrimSpace(c.Name))
// 	c.UpdatedDt = time.Now()

// 	if c.Stat != 0 && c.Stat != 1 {
// 		c.Stat = 1 // default ke Active
// 	}
// 	// if err := db.Delete(&model.Category{}, id).Error; err != nil {
// 	// 			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
// 	// 		}
// 	// Ambil data lama
// 	var old Category
// 	if err := tx.Unscoped().First(&old, c.ID).Error; err != nil {
// 		return nil // skip silently
// 	}

// 	// Bandingkan field tertentu
// 	changes := compareFields(&old, c, []string{},
// 		[]string{"ID", "CreatedDt", "CreatedBy", "UpdatedDt", "UpdatedBy", "LastSyncDt"})

// 	// Simpan perubahan ke log DB
// 	for field, val := range changes {
// 		log := LogCategory{
// 			PKID:      c.ID,
// 			FieldName: field,
// 			CreatedBy: c.CreatedBy,
// 			CreatedDt: c.CreatedDt,
// 			OldVal:    fmt.Sprintf("%v", val[0]),
// 			NewVal:    fmt.Sprintf("%v", val[1]),
// 		}
// 		log.SaveLogCategory()
// 	}

// 	return nil
// }
