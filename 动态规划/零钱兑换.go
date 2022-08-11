package main

import "fmt"

// 完全背包问题：

//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
//你可以认为每种硬币的数量是无限的。


//自下而上的方式进行思考 F(i) 为组成金额 i 所需最少的硬币数量,假设在计算 F(i) 之前，我们已经计算出 F(0)-F(i-1) 的答案。 则 F(i) 对应的转移方程应为
// f(i) = min f(i-c(j)) +1           j=0...n-1
// c(j) 代表第j枚硬币的面值，

func main() {
	fmt.Println(coinChange([]int{1, 2, 5},11))
}

func coinChange(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, max)
	for i:= range dp {
		dp[i] = max
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp [i] =  min(dp[i], dp[i-coins[j]] +1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}else {
		return dp[amount]
	}
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}


