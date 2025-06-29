package main

// func Sum(nums [5]int) (total int) {
// 	for i := 0; i < 5; i++ {
// 		total += nums[i]
// 	}

// 	return
// }

// introducing range
func Sum(nums []int) (total int) {
	for _, num := range nums { // range allows iteration over an array
		total += num
	}
	return
}

// arrays in go must have a fixed length
// to have varying length go has what we call slices

func SumAll(arrToSum ...[]int) []int {
	lenofArg := len(arrToSum)     // essentially defines how many args of arr are in function
	sums := make([]int, lenofArg) // make function declares a slice of a defined length

	for i, arrs := range arrToSum {
		sums[i] = Sum(arrs)
	}
	return sums
}
