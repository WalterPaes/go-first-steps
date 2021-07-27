package math

import "testing"

func TestAdd(t *testing.T) {
	result := add(1, 2)
	expected := 3

	if result != expected {
		t.Errorf("Was expected '%d', but got '%d'", expected, result)
	}
}
