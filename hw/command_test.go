package hw

import "testing"

func TestExecute(t *testing.T) {
	exp := "46.54"
	got, err := execute("echo", "46540")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if got != exp {
		t.Errorf("Expected: %v, got: %v.", exp, got)
	}
}
