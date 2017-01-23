package hw

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Run("temp=49.8'C=>49.8", parseSuccess("temp=49.8'C", 49.8))
	t.Run("46.540=>46.54", parseSuccess("46.540", 46.54))
	t.Run("1=>1", parseSuccess("1", 1))
	t.Run("1.0=>1", parseSuccess("1.0", 1))
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
