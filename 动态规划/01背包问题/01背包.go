package main

import "fmt"

/*
方法1 二维数组
dp[i][j] 表示从下标[0-i]的物品里任意取，放进容量为j的背包，价值总和最大时多少？

递归公式： dp[i][j] = max(dp[i - 1][j], dp[i - 1][j - weight[i]] + value[i]);

*/

// 二维dp
func test_2_wei_bag_problem1(weight []int, value []int, bagWeight int) int {

	dp := make([][]int, len(weight))

	for r := range dp {
		dp[r] = make([]int, bagWeight+1) //  bagWeight = 0 没有任何意义，只是动态规划的递推使用的。所以bagWight+1
	}
	// 初始首行，第0个物品，背包重量为 0，1，2，3，4 ...bagWeight
	for j := 1; j <= bagWeight && j >= weight[0]; j++ {
		dp[0][j] = value[0]
	}
	// 初始化首列,  背包为0 ，所以不可能装物品，初始化为0


	// dp过程,先物品，后重量
	for i := 1; i < len(weight); i++ { // 第一个物品已经算过了
		for j := 0; j <= bagWeight; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = maxInt(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	return dp[len(weight)-1][bagWeight]

}

func test_2_wei_bag_problem2(weight []int, value []int, bagWeight int) int {

	// 定义dp数组 dp[j]背包重量为j最大价值
	dp := make([]int, bagWeight+1)

	for i := 0; i < len(weight); i++ {
		for j := bagWeight; j >= weight[i]; j-- {
			dp[j] = maxInt(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	return dp[bagWeight]
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	fmt.Println(test_2_wei_bag_problem1(weight, value, 4))
}
