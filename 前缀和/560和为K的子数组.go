package main

//给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,1,1], k = 2
//输出：2
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3], k = 3
//输出：2

func subarraySum(nums []int, k int) int {
	preSum := 0
	cnt := 0
	//pre[i] - pre[j] == k
	//pre[j] = pre[i] - k, 使用mp记录pre[j]的个数
	mp := map[int]int{0: 1}

	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		if mp[preSum-k] > 0 {
			cnt += mp[preSum-k]
		}
		mp[preSum]++
	}
	return cnt
}
