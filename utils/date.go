package utils

import (
	"strings"
	"time"
)

func NewDate(s string) string {
	short := "2006-01-02"

	if strings.Index("all", s) == 0 {
		return ""
	} else if strings.Index("now", s) == 0 {
		return time.Now().Format(time.RFC3339)
	} else if strings.Index("today", s) == 0 {
		return time.Now().Format(short)
	} else if strings.Index("yesterday", s) == 0 {
		return time.Now().Add(-time.Hour * 24).Format(short)
	}

	return s
}
