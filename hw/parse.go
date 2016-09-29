package hw

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func parseTemperature(temp string) (string, error) {
	t, err := strconv.ParseFloat(strings.TrimSpace(string(temp)), 64)
	if err != nil {
		return "", errors.Wrap(err, "could not parse temperature")
	}
	return strconv.FormatFloat(t/1000, 'f', 2, 64), nil
}
