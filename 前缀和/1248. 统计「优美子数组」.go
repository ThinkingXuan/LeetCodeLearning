package main

//给你一个整数数组nums 和一个整数 k。如果某个连续子数组中恰好有 k 个奇数数字，我们就认为这个子数组是「优美子数组」。
//
//请返回这个数组中 「优美子数组」 的数目。

//https://leetcode.cn/problems/count-number-of-nice-subarrays/solution/hua-dong-chuang-kou-qian-zhui-he-bi-xu-miao-dong-b/
func numberOfSubarrays(nums []int, k int) int {
	//key是前缀和（即当前奇数的个数），值是前缀和的个数。
	fixFixCnt := make(map[int]int)
	fixFixCnt[0] = 1
	// 遍历原数组，计算当前的前缀和，统计到 prefixCnt 数组中，
	// 并且在 res 中累加上与当前前缀和差值为 k 的前缀和的个数。

	res := 0
	preSum := 0
	for i := 0; i < len(nums); i++ {
		// 奇数&1 = 1 偶数&1=0
		preSum += nums[i] & 1
		if fixFixCnt[preSum-k] > 0 {
			res += fixFixCnt[preSum-k]
		}
		fixFixCnt[preSum]++
	}
	return res
}
