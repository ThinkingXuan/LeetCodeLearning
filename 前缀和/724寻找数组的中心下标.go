package main

//给你一个整数数组 nums ，请计算数组的 中心下标 。
//
// 数组 中心下标 是数组的一个下标，其左侧所有元素相加的和等于右侧所有元素相加的和。
//
// 如果中心下标位于数组最左端，那么左侧数之和视为 0 ，因为在下标的左侧不存在元素。这一点对于中心下标位于数组最右端同样适用。
//
// 如果数组有多个中心下标，应该返回 最靠近左边 的那一个。如果数组不存在中心下标，返回 -1 。
//
//
//
// 示例 1：
//输入：nums = [1, 7, 3, 6, 5, 6]
//输出：3
//解释：
//中心下标是 3 。
//左侧数之和 sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11 ，
//右侧数之和 sum = nums[4] + nums[5] = 5 + 6 = 11 ，二者相等。

// Related Topics 数组 前缀和 👍 462 👎 0

func pivotIndex(nums []int) int {
	preSum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		// 从0开始 [0: i]  ans = preSum[i+1] - preSum[0]
		// 从0开始 [i: n-1] ans = preSum[n] - preSum[i]
		if preSum[i+1]-preSum[0] == preSum[n]-preSum[i] {
			return i
		}
	}
	return -1
}
