package main

//给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums)*2)

	for i := 0; i < len(res); i++ {
		res[i] = -1
	}

	tmp := make([]int, len(nums))
	copy(tmp, nums)

	for i := 0; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
	}

	var st []int
	for i := 0; i < len(tmp); i++ {

		for len(st) > 0 && tmp[st[len(st)-1]] < tmp[i] {
			top := st[len(st)-1]
			res[top] = tmp[i]

			st = st[:len(st)-1]
		}

		st = append(st, i)
	}
	return res[:len(nums)]
}
