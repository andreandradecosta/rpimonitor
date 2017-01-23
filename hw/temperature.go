// +build !arm

package hw

// GetTemperature returns the host temperature or '-' if not available.
func GetTemperature() (float64, error) {
	return 0, nil
}
