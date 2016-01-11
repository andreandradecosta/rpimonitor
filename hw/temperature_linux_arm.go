// +build linux
// +build arm

package hw

func GetTemperature() (string, error) {
	return execute("/opt/vc/bin/vcgencmd", "measure_temp")
}
