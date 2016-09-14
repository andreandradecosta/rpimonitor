package device

import "testing"

func TestParse(t *testing.T) {
	t.Run("46540=>46.54", parseSuccess("46540", "46.54"))
	t.Run("-46540=>-46.54", parseSuccess("-46540", "-46.54"))
	t.Run("xxx=>Error", parseError("xxx"))
}

func parseSuccess(text, exp string) func(*testing.T) {
	return func(t *testing.T) {
		got, err := parseTemperature(text)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
		if got != exp {
			t.Errorf("Expected %v, got %v", exp, got)
		}
	}
}

func parseError(text string) func(*testing.T) {
	return func(t *testing.T) {
		if _, err := parseTemperature(text); err == nil {
			t.Errorf("Expected error, got nil.")
		}
	}
}
