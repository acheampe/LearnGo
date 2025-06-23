package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// Document function
func ExampleRepeat() {
	fmt.Println(Repeat("a", 5))
	// Output: aaaaa
}

// benchmark is a first class feature of go
// it provides assess to a loop func
// measures performance of code
// use when comparing two ways to solve a problem
// optimizing critical code path
// test impact of allocation and memory use
// use benchmem to track memory usage: go test -bench=. -benchmem
func BenchmarkRepeat(b *testing.B) {
	// ...set up...
	for b.Loop() {
		//...code to measure (only the body of the loop is measure)...
		Repeat("a", 6)
	}
	//...cleanup...
}
