package hw

import "testing"

func TestExecute(t *testing.T) {
	exp := "51'C"
	got, err := execute("echo", "temp=51'C")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if got != exp {
		t.Errorf("Expected: %v, got: %v.", exp, got)
	}
}
