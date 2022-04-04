package validation

import "time"

func Weekday(date time.Time) time.Time {
	if date.Weekday().String() == "Saturday" {
		date = date.AddDate(0, 0, -1)
		return date
	} else if date.Weekday().String() == "Sunday" {
		date = date.AddDate(0, 0, -2)
		return date
	}
	return date
}
