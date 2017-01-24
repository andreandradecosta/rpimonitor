package hw

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

//parseTemperature converts the string temp from "50306" to 50.306
func parseTemperature(temp string) (float64, error) {
	t, err := strconv.ParseFloat(strings.TrimSpace(temp), 64)
	if err != nil {
		return 0, errors.Wrapf(err, "Could not parse temperature from: %s", temp)
	}
	return t / 1000, nil
}
