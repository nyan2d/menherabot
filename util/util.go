package util

import (
	"fmt"
	"strconv"
	"time"
)

func TimeToInt(t time.Time) int {
	strtime := fmt.Sprintf("%v%v%v", t.Year(), t.Month(), t.Day())
	i, _ := strconv.Atoi(strtime)
	return i
}
