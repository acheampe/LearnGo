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
