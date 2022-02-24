package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	result := Add(2, 2)
	expected := 4

	if result != expected {
		t.Errorf("result was: %d, but expected: %d", result, expected)
	}

}

func ExampleAdd() {
	sum := Add(3, 3)

	fmt.Println(sum)
	// Output: 6
}
