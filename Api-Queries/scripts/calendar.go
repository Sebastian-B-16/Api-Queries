package scripts

import "time"

func Calendar(year int, month int, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
