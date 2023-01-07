package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	max := nums[0]

	// for i := 1; i < len(nums); i++ {
	// 	if nums[i] > max {
	// 		max = nums[i]
	// 	}
	// }

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	fmt.Printf("%v", max)
}
