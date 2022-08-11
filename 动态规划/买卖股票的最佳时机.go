package main

import "math"

//给定一个数组 prices ，它的第i 个元素prices[i] 表示一支给定股票第 i 天的价格。
//
//你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
//
//返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

func main() {

}

// 暴力法，
// 计算出所有利润，然后比较大小即可
func maxProfit(prices []int) int {

	maxProfit := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > maxProfit {
				maxProfit = prices[j] - prices[j]
			}
		}
	}
	return maxProfit
}

// 一次遍历法
// 找到最低价格，然后计算利润
func maxProfit2(prices []int) int {

	minProfit := math.MaxInt32
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minProfit {
			minProfit = prices[i]
		} else if prices[i]-minProfit > maxProfit {
			maxProfit = prices[i] - minProfit
		}
	}

	return maxProfit
}

// 动态规划法

