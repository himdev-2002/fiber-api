package structs

import (
	"github.com/phuslu/log"
)

type Logger struct {
	Console log.Logger
	Error   log.Logger
	Data    log.Logger
}
