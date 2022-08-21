package main

import "fmt"

/*
差分数组：
nums为原数组， 差分数组就是遵循 diff[i] = nums[i] - nums[i-1]
差分数组还原成原数组, 则需要ret[i] = ret[i-1] + nums[i]

使用差分数组好处？
如果需要修改数组left到right得值，不需要遍历，只需要将diff[left] -= data 且 diff[right+1] += data即可
*/

func main() {
	nums := []int{1, 2, 6, 7, 9}
	diff := getDiff(nums)

	increase(1, 5, diff, 2)
	fmt.Println(getNums(diff))

	nums2 := []int{0, 0, 0, 0, 0}
	fmt.Println(numsIncrease(nums2, [][]int{{1, 3, 2}, {2, 4, 3}, {0, 2, -2}}))

}

func getDiff(nums []int) []int {
	diff := make([]int, len(nums))
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return diff
}

// 差分数组构造原数组
func getNums(diff []int) []int {
	nums := make([]int, len(diff))
	nums[0] = diff[0]
	for i := 1; i < len(diff); i++ {
		nums[i] = nums[i-1] + diff[i]
	}
	return nums
}

// 更改原数组的值（left到right）下标从0开始。
func increase(left, right int, diff []int, num int) {
	diff[left] += num
	if right+1 < len(diff) {
		diff[right+1] -= num
	}
}

// LeetCode 370 区间加法 会员题
//假设你有一个长度为n的数组，初始情况下所有的数字均为0，你将会被分给出k个更新的操作。
//其中， 每个操作会被表示为一个三元组：[startIndex, endIndex, inc], 你需要将子数组
// A[startIndex, ...endIndex]包括（startIndex和EndIndex）添加inc，请返回k次操作的数组

func numsIncrease(nums []int, a [][]int) []int {

	diff := make([]int, len(nums))
	diff[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}

	for i := 0; i < len(a); i++ {
		start, end, inc := a[i][0], a[i][1], a[i][2]

		diff[start] += inc
		if end+1 < len(diff) {
			diff[end+1] -= inc
		}
	}

	// 还原
	nums[0] = diff[0]
	for i := 1; i < len(diff); i++ {
		nums[i] = nums[i-1] + diff[i]
	}
	return nums
}

// 1109 航班预定问题
// 这里有 n 个航班，它们分别从 1 到 n 进行编号。
//
// 有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi] 意味着在从
//firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。
//
// 请你返回一个长度为 n 的数组 answer，里面的元素是每个航班预定的座位总数。
//

func corpFlightBookings(bookings [][]int, n int) []int {
	diff := make([]int, n)

	for _, v := range bookings {
		start, end, inc := v[0]-1, v[1]-1, v[2]

		diff[start] += inc
		if end+1 < n {
			diff[end+1] -= inc
		}
	}

	res := make([]int, n)
	res[0] = diff[0]

	for i := 1; i < n; i++ {
		res[i] = res[i-1] + diff[i]
	}
	return res
}

// 1094 拼车
//车上最初有 capacity 个空座位。车 只能 向一个方向行驶（也就是说，不允许掉头或改变方向）
//
// 给定整数 capacity 和一个数组 trips , trip[i] = [numPassengersi, fromi, toi] 表示第 i 次旅行有
// numPassengersi 乘客，接他们和放他们的位置分别是 fromi 和 toi 。这些位置是从汽车的初始位置向东的公里数。
//
// 当且仅当你可以在所有给定的行程中接送所有乘客时，返回 true，否则请返回 false。

/**
看准边界条件
i从0开始，
left, right 边界 不包括right，因为乘客有下车的
*/

func carPooling(trips [][]int, capacity int) bool {
	diff := make([]int, 10001)

	for _, v := range trips {
		//即乘客在车上的区间是 [trip[1], trip[2] - 1]
		left, right, inc := v[1], v[2]-1, v[0]

		diff[left] += inc
		if right+1 < len(diff) {
			diff[right+1] -= inc
		}
	}

	res := make([]int, 10001)
	res[0] = diff[0]

	for i := 1; i < len(res); i++ {
		res[i] = res[i-1] + diff[i]
	}

	for i := 0; i < len(res); i++ {
		if res[i] > capacity {
			return false
		}
	}
	return true
}
