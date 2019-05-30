package golog

import (
	"github.com/rs/zerolog"
)

func init() {
	zerolog.TimestampFieldName = "time"
	zerolog.LevelFieldName = "level"
	zerolog.MessageFieldName = "msg"
}
