package scripts

import (
	"Api-Queries/validation"
	"time"
)

type YearlyPayDays struct {
	Date string `json:"date"`
}

func YearlyPay(day int) []YearlyPayDays {
	currentTime := time.Now()
	currentYear := currentTime.Year()
	var output []YearlyPayDays

	for i := int(currentTime.Month()); i <= 12; i++ {
		paydayDate := Calendar(currentYear, i, day)
		paydayDate = validation.Weekday(paydayDate)
		tempOutput := YearlyPayDays{Date: paydayDate.Format("02-01-2006")}
		output = append(output, tempOutput)
	}
	return output
}
