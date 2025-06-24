package services

import (
	"fmt"
	"os"
	"strconv"
	"him/fiber-api/core/structs"

	"github.com/phuslu/log"
)

var logger = structs.Logger{
	Console: log.Logger{
		TimeFormat: "15:04:05",
		Caller:     1,
		Writer: &log.ConsoleWriter{
			ColorOutput:    true,
			EndWithMessage: true,
		},
	},
	Error: log.Logger{
		Level: log.ErrorLevel,
		Writer: &log.FileWriter{
			Filename:   "logs/error.log",
			MaxSize:    50 * 1024 * 1024,
			MaxBackups: 7,
			LocalTime:  false,
		},
	},
	Data: log.Logger{
		Level: log.InfoLevel,
		Writer: &log.FileWriter{
			Filename:   "logs/data.log",
			MaxSize:    50 * 1024 * 1024,
			MaxBackups: 7,
			LocalTime:  false,
		},
	},
}

func DebugLog() (*log.Logger, error) {
	if level, err := strconv.Atoi(os.Getenv("LOG_LEVEL")); err == nil && level <= int(log.DebugLevel) {
		return &logger.Console, nil
	} else {
		return nil, fmt.Errorf("lower log level")
	}
}

func InfoLog() (*log.Logger, error) {
	if level, err := strconv.Atoi(os.Getenv("LOG_LEVEL")); err == nil && level <= int(log.InfoLevel) {
		return &logger.Console, nil
	} else {
		return nil, fmt.Errorf("lower log level")
	}
}

func ErrorLog() (*log.Logger, error) {
	if level, err := strconv.Atoi(os.Getenv("LOG_LEVEL")); err == nil && level <= int(log.ErrorLevel) {
		return &logger.Error, nil
	} else {
		return nil, fmt.Errorf("lower log level")
	}
}

func DataLog(msg string) {
	logger.Data.Log().Msgf(msg)
}
