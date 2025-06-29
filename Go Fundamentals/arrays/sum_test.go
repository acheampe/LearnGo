package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers) // %v prints default format
		}
	})

}

func ExampleSum() {
	fmt.Println(Sum([]int{1, 2, 3, 4, 5}))
	// Output: 15
}

func TestSumAll(t *testing.T) {
	t.Run("Sum all arr into respective index of new output array", func(t *testing.T) {
		num1 := []int{1, 2, 3}
		num2 := []int{2, 3}

		got := SumAll(num1, num2)
		want := []int{6, 5}

		// cannot use equality operators with slices in go so we
		// can use reflect.DeepEqual instead
		if !reflect.DeepEqual(got, want) { // NOTE: reflect is not type safe
			// it will compile comparisons between string and int so its use must be with care
			t.Errorf("got %v want %v", got, want)
		}
	})
}
