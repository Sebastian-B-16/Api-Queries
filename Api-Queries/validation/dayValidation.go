package validation

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
)

var queryCheck = regexp.MustCompile(`^[0-9]+$`)

func Query(r *http.Request) (int, int) {
	query, ok := r.URL.Query()["pay_day"]
	isValid := queryCheck.MatchString(query[0])
	if !isValid {
		return http.StatusBadRequest, 0
	}
	payDay, err := strconv.Atoi(query[0])
	if err != nil {
		log.Print("Error converting the query into Int,", err)
		return http.StatusBadRequest, 0
	}

	if payDay < 1 || payDay > 31 || !ok {
		return http.StatusBadRequest, 0
	}
	return http.StatusOK, payDay
}
