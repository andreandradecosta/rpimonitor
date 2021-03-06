// +build linux
// +build arm

package hw

import (
	"os/exec"

	"github.com/pkg/errors"
)

func GetTemperature() (float64, error) {
	cmd := exec.Command("cat", "/sys/class/thermal/thermal_zone0/temp")
	output, err := cmd.Output()
	if err != nil {
		return 0, errors.Wrap(err, "could not read temperature")
	}
	return parseTemperature(string(output))
}
