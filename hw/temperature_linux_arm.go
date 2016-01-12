// +build linux
// +build arm

package hw

func GetTemperature() (string, error) {
	return execute("cat", "/sys/class/thermal/thermal_zone0/temp")
}
