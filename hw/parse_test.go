package hw

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Run("50306=>50.306", parseSuccess("50306", 50.306))
	t.Run("46540=>46.54", parseSuccess("46540", 46.54))
	t.Run("46540CR=>46.54", parseSuccess("46540\n", 46.54))
	t.Run("1=>1", parseSuccess("1", 0.001))
	t.Run("xxx=>Error", parseError("xxx"))
}

func parseSuccess(text string, exp float64) func(*testing.T) {
	return func(t *testing.T) {
		assert := assert.New(t)
		actual, err := parseTemperature(text)
		if assert.NoError(err) {
			assert.Equal(exp, actual)
		}
	}
}

func parseError(text string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := parseTemperature(text)
		assert.Error(t, err)
	}
}
