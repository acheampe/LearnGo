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

// Good for unknown size but more memory allocation
// func SumAll(arrToSum ...[]int) []int {
// 	var sums []int

// 	for _, arr := range arrToSum {
// 		sums = append(sums, Sum(arr)) // appending a new value to sum arr
// 	}
// 	return sums
// }

// func SumAllTails(arrToSum ...[]int) []int {
// 	argLen := len(arrToSum)
// 	tailSum := make([]int, argLen)
// 	for i, arr := range arrToSum {
// 		tailSum[i] = Sum(arr) - arr[0]
// 	}
// 	return tailSum
// }

func SumAllTails(arrToSum ...[]int) []int {
	argLen := len(arrToSum)
	tailSum := make([]int, argLen)
	for i, arr := range arrToSum {
		if len(arr) == 0 {
			tailSum[i] = 0
		} else {
			tailSum[i] = Sum(arr[1:])
		}
	}
	return tailSum
}
