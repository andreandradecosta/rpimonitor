package hw

import (
	"os/exec"
	"strconv"
	"strings"
)

func execute(command, arg string) (string, error) {
	cmd := exec.Command(command, arg)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	t, err := strconv.ParseFloat(strings.TrimSpace(string(output)), 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatFloat(t/1000, 'f', 2, 64), nil
}
