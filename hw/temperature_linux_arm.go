// +build linux
// +build arm

package hw

import (
	"os/exec"

	"github.com/pkg/errors"
)

func GetTemperature() (float64, error) {
	cmd := exec.Command("/opt/vc/bin/vcgencmd", "measure_temp")
	output, err := cmd.Output()
	if err != nil {
		return "", errors.Wrap(err, "could not read temperature")
	}
	return parseTemperature(string(output))
}
