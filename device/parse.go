package device

import (
	"strconv"
	"strings"
)

func parseTemperature(temp string) (string, error) {
	t, err := strconv.ParseFloat(strings.TrimSpace(string(temp)), 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatFloat(t/1000, 'f', 2, 64), nil
}
