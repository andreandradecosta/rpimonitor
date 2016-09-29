// +build linux
// +build arm

package hw

import (
	"os/exec"

	"github.com/pkg/errors"
)

func GetTemperature() (string, error) {
	cmd := exec.Command("cat", "/sys/class/thermal/thermal_zone0/temp")
	output, err := cmd.Output()
	if err != nil {
		return "", errors.Wrap(err, "could not read temperature")
	}
	return parseTemperature(string(output))
}
