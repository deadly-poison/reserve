package util

import (
	"fmt"
	"time"
)

func CheckTime(timeStr string) bool {
	startTime, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr , time.Local)
	if err != nil {
		fmt.Println()
	}
	return time.Now().After(startTime)
}
