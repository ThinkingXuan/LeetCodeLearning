package main

//你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的
//房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，今晚能够偷窃到的最高金额。

// 1 <= nums.length <= 100
// 0 <= nums[i] <= 1000
//
//
// Related Topics 数组 动态规划 👍 1151 👎 0

//dp[i]
// dp[i] = max(dp[i-2]+nums[i], dp[i-1])
func rob2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	// 考虑首尾
	// 1. 不包含首，包含尾
	// 2. 包含首，不包含尾
	result1 := robRange(nums, 1, len(nums)-1)
	result2 := robRange(nums, 0, len(nums)-2)

	return max(result1, result2)
}

func robRange(nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}

	dp := make([]int, len(nums))
	dp[start] = nums[start]
	dp[start+1] = max(nums[start], nums[start+1])

	for i := start + 2; i <= end; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[end]
}
