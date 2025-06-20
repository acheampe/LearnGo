package integers // instead of main package this will group functions for working with integers

import "testing"

func TestAdder(t *testing.T) {
	var sum int
	sum = Add(2, 2)
	expected := 4

	// %q if we want it to print a string, %d if we want to print and integer
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
