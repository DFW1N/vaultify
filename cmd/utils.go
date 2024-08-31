package cmd

import (
	"time"
)

func parseDate(date string) (time.Time, error) {
	if date == "" {
		return time.Time{}, nil
	}
	return time.Parse("2006-01-02", date)
}
