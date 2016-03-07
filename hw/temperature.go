// +build !arm

package hw

//GetTemperature returns the host temperature or '-' if not available.
func GetTemperature() (string, error) {
	return "-", nil
}
