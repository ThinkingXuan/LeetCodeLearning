package main

/*给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

子数组 是数组中的一个连续部分。*/
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	res := nums[0]
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		if sum < nums[i] {
			sum = nums[i]
		}
		res = max(res, sum)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
