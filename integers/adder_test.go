package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("sum was %d but expected it to be %d", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
