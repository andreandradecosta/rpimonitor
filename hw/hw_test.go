package hw

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetdata(t *testing.T) {
	t.Run("Return Data", returnData)
	t.Run("Return Error Message", returnError)
}

func returnData(t *testing.T) {
	assert.Equal(t, "data", getData("data", nil))
}

func returnError(t *testing.T) {
	assert.Equal(t, "error message", getData("data", errors.New("error message")))
}
