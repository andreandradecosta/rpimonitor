package hw

import (
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

//parseTemperature extract numeric temperature from a string like temp=49.8'C
func parseTemperature(temp string) (float64, error) {
	r, _ := regexp.Compile("([0-9]*\\.[0-9]+|[0-9]+)")
	tStr := r.FindString(temp)
	if len(tStr) == 0 {
		return 0, errors.Errorf("Could not extract temperature from: %s", temp)
	}
	t, _ := strconv.ParseFloat(tStr, 64)
	return t, nil
}
