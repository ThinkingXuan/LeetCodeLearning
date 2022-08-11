package main

import "fmt"

// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
// 分析：记 f(i) 为第i坐标的前i个连续数值最大和，
// 所以题目结果为      max{f(i)}  0<=i<=n
// 那么就是求f(i)的问题了，  f(i) 的取值与 num[i] 和f(i-1)+num[i] 有关，我们希望获取一个比较大的数值。
// 所以动态转移方程：
//   f(i) = max{ f(i-1)+ nums[i], nums[i]}
//   边界条件  f(0) = nums[0]

func main() {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}

func maxSubArray(nums []int) int {
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if  nums[i] + nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

