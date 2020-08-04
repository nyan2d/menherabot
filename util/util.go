package util

import (
	"time"
)

func TimeToInt(t time.Time) int64 {
	return t.Truncate(time.Hour * 24).Unix()
}
