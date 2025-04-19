package main

/*
496.下一个更大元素 I

给你两个 没有重复元素 的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。

请你找出 nums1 中每个元素在 nums2 中的下一个比其大的值。

nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出 -1 。
*/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))

	for i := range res {
		res[i] = -1
	}

	mp := make(map[int]int)
	for i, v := range nums1 {
		mp[v] = i
	}

	var st []int
	for i := 0; i < len(nums2); i++ {

		for len(st) > 0 && nums2[st[len(st)-1]] < nums2[i] {
			top := st[len(st)-1]
			if v, ok := mp[nums2[top]]; ok {
				res[v] = nums2[i]
			}
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return res
}
