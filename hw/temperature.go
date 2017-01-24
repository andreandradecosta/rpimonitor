// +build !arm

package hw

// GetTemperature returns the host temperature or 0 if not available.
func GetTemperature() (float64, error) {
	return 0, nil
}
