package hw

import (
	"os/exec"
	"strings"
)

func execute(command, arg string) (string, error) {
	cmd := exec.Command(command, arg)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	val := strings.TrimSpace(strings.Split(string(output), "=")[1])
	return val, nil
}
