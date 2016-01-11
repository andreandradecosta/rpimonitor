// +build !arm

package hw

func GetTemperature() (string, error) {
	return "-", nil
}
