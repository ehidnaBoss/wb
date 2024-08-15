package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 1, 6, 5, 3, 4}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Println(nums)
}
