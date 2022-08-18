package main

// 模板代码.
func findSubArray(nums []int) int {
	n := len(nums)

	left, right := 0, 0 // 双指针，表示当前遍历的区间[left, right]，闭区间
	sum := 0            //用于统计 子数组/子区间 是否有效，根据题目可能会改成求和/计数
	res := 0            // 保存最大的满足题目要求的 子数组/子串 长度

	for right < n { // 当右边的指针没有搜索到 数组/字符串 的结尾
		sum += nums[right] // 增加当前右边指针的数字/字符的求和/计数

		for true { //区间[left, right]不符合题意
			sum -= nums[left] //移动左指针前需要从counter中减少left位置字符的求和/计数
			left += 1         //真正的移动左指针，注意不能跟上面一行代码写反
		}
		res = max(res, right-left+1) // 需要更新结果
		right += 1                   // 移动右指针，去探索新的区间
	}
	return res
}
