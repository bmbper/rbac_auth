package base

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

func Guid() string {
	return uuid.New().String()
}
func SimpleGuid() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func NowTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
