package utils

import "time"

func GetDateTimeNot() string {
	return time.Now().Format("2006010215:04:05")
}

func GetDate() string {
	return time.Now().Format("2006-01-02")
}

func GetBasicDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetNowTimesTamp() string {
	return time.Now().Format("20060102150405")
}

func GetDateTimeBeforeHours(hour int) string {
	return time.Now().Add(-time.Hour * time.Duration(hour)).Format("2006-01-02 15:04:05")
}

func GetDateBeforeDays(days int) string {
	return time.Now().Add(-time.Hour * time.Duration(days) * 24).Format("2006-01-02")
}

func GetDateTimeBeforeDays(days int) string {
	return time.Now().Add(-time.Hour * time.Duration(days) * 24).Format("2006-01-02 15:04:05")
}
