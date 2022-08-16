package main

func main() {

}
func largestRectangleArea(heights []int) int {
	// 递增栈，找左右比栈顶mid小的。
	res := 0

	// 第一个和最后一个也要计算
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)

	var st []int
	st = append(st, 0)
	for i := 1; i < len(heights); i++ {
		for len(st) > 0 && heights[st[len(st)-1]] > heights[i] {
			mid := st[len(st)-1]
			st = st[:len(st)-1]

			// 左边比mid大的left， 右边比mid大的 i
			left := st[len(st)-1]
			w := i - left - 1
			h := heights[mid]

			res = max(res, w*h)
		}
		st = append(st, i)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
