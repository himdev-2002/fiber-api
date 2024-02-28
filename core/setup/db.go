package setup

import (
	"os"
	"tde/fiber-api/core/models"
	"tde/fiber-api/core/services"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDBSys() {
	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("Connecting to Sys DB...")
	}
	var err error
	if os.Getenv("DB_SYS_DRIVER") == "sqlite" {
		services.DBCore, err = gorm.Open(sqlite.Open(os.Getenv("DB_SYS_URL")), &gorm.Config{})
	}

	if err != nil {
		if log, err2 := services.ErrorLog(); err2 == nil {
			log.Log().Msgf("Failed connect to Sys DB")
		}
		panic("Failed connect to Sys DB")
	} else {
		if log, err := services.InfoLog(); err == nil {
			log.Info().Msgf("CONNECTED to Sys DB!")
		}
		migrateDBSys()
	}
}

func migrateDBSys() {
	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("Migrate Sys DB...")
	}
	services.DBCore.Table("m_user").AutoMigrate(&models.User{})
}

func DisconnectDBSys() {
	if log, err := services.DebugLog(); err == nil {
		log.Debug().Msgf("Disconnecting from Sys DB...")
	}
	db, err := services.DBCore.DB()
	db.Close()

	if err != nil {
		if log, err2 := services.ErrorLog(); err2 == nil {
			log.Log().Msgf("Failed disconnect from Sys DB")
		}
		panic("Failed disconnect from Sys DB")
	} else {
		if log, err := services.InfoLog(); err == nil {
			log.Info().Msgf("DISCONNECTED from Sys DB!")
		}
	}
}
