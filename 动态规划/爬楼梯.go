package main

import "fmt"

//假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

//分析
//一个阶梯  有1种   记为f(1)
//二个阶梯  有2种   记为f(2)
//三个阶梯  有3种   记为f(3) = f(1) + f(2)
//n个阶梯        有 f(n) = f(n-1) + f(n-2)

//1 2 3 5 8 13 21 34 55  89
func main() {
	fmt.Println(climbStairs(10))
	fmt.Println(climbStairs2(10))
	fmt.Println(climbStairs3(10))
}

func climbStairs(n int) int {

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

// 非递归实现
func climbStairs2(n int) int {

	if n <= 2 {
		return n
	}
	dp := make([]int,n+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 非递归实现，优化内存使用
func climbStairs3(n int) int {

	if n <= 2 {
		return n
	}
	dp := make([]int,4)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[3] = dp[1] +dp[2]
		dp[1] = dp[2]
		dp[2] = dp[3]
	}
	return dp[3]
}

