package scripts

import (
	"Api-Queries/validation"
	"time"
)

type PayDay struct {
	NextPayDayIn int `json:"next_payday"`
}

func NextPay(day int) []PayDay {
	currentTime := time.Now()
	currentYear := currentTime.Year()
	currentMonth := currentTime.Month()
	date := Calendar(currentYear, int(currentMonth), day)
	if date.Day()-currentTime.Day() < 0 {
		date = Calendar(currentYear, int(currentMonth+1), day)
	}
	date = validation.Weekday(date)
	diff := date.Sub(currentTime)
	outcome := int(diff.Hours() / 24)
	var output []PayDay
	tempOutput := PayDay{NextPayDayIn: outcome}
	output = append(output, tempOutput)
	return output

}
