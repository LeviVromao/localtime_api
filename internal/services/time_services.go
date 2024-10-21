package services

import (
	"fmt"
	"time"
)

var cityTimeZones = map[string]string{
	"SP":  "America/New_York",
	"LON": "Europe/London",
	"BER": "Europe/Berlin",
	"TYO": "Asia/Tokyo",
	"CCU": "Asia/Kolkata",
	"JNB": "Africa/Johannesburg",
	"SYD": "Australia/Sydney",
	"AKL": "Pacific/Auckland",
}

func getCityTimeZone(city string) (string, bool) {
	timezone, exists := cityTimeZones[city]
	return timezone, exists
}

func GetLocalTime(city string) (string, error) {
	timezone, exists := getCityTimeZone(city)
	if !exists {
		return "", fmt.Errorf("City or Timezone not found")
	}
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return "", err
	}

	return time.Now().In(loc).Format("2006-01-02 15:04:05"), nil
}
