package main

import "sort"

func main() {
	nums := []int{2, 34, 35, 1, 34, 12, 243}
	// 不稳定
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	// 稳定
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

}
