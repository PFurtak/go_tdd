package integers

import "testing"

func TestAdder(t *testing.T) {
	result := Add(2, 2)
	expected := 4

	if result != expected {
		t.Errorf("result was: %d, but expected: %d", result, expected)
	}

}
