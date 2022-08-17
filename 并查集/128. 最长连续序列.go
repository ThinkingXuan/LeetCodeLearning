package main

/*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为O(n) 的算法解决此问题。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-consecutive-sequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func longestConsecutive(nums []int) int {
	initialize()
	mp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := mp[nums[i]]; ok {
			continue
		}

		if _, ok := mp[nums[i]-1]; ok {
			join(i, mp[nums[i]-1])
		}

		if _, ok := mp[nums[i]+1]; ok {
			join(i, mp[nums[i]+1])
		}
		mp[nums[i]] = i
	}
	return connectMaxSize(len(nums))
}
