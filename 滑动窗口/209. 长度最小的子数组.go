package main

import "math"

/*给定一个含有个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/minimum-size-subarray-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func minSubArrayLen(target int, nums []int) int {
	res := math.MaxInt32
	sum := 0
	j := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]

		for sum >= target {
			res = min(res, i-j+1)
			sum -= nums[j]
			j++

		}
	}
	if res == math.MaxInt32 {
		return 0
	}

	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
