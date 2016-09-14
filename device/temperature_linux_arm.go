// +build linux
// +build arm

package device

import "os/exec"

func GetTemperature() (string, error) {
	cmd := exec.Command("cat", "/sys/class/thermal/thermal_zone0/temp")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return parseTemperature(output)
}
