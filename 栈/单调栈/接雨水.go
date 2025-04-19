package main

//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
// 经典题目：
// 方法一：双指针
// 方法二：动态规划
// 方法三：单调栈

func main() {

}

func trap(height []int) int {
	// 递减栈
	var st []int
	sum := 0

	// 第一个不计算，先入栈
	st = append(st, 0)

	for i := 1; i < len(height); i++ {

		for len(st) > 0 && height[st[len(st)-1]] < height[i] {
			// 中间的
			mid := st[len(st)-1]
			st = st[:len(st)-1]

			if len(st) > 0 {
				// st[len(st)-1] 比中间大的左边第一个， i 别中间大的右边第一个
				h := min(height[st[len(st)-1]], height[i]) - height[mid]
				w := i - st[len(st)-1] - 1 // 宽度
				sum += h * w
			}
		}
		st = append(st, i)
	}
	return sum
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
