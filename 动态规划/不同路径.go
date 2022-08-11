package main

//一个机器人位于一个 m x n网格的左上角 （起始点在下图中标记为 “Start” ）。
//
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//
//问总共有多少条不同的路径？

// 解析：设 f(i,j) 表示从左上角走到(i,j)的路径数量， i,j的范围分别时[0,m)和[0,n)
// 我们每一步只能从向下或者向右移动一步 ,此我们可以写出动态规划转移方程：
//  f(i,j) = f(i-1,j) + f(i,j-1)
//  初始条件  f(0,0) = 1

func main() {
	println(uniquePaths(3,7))
}

//我们可以将所有的 f(0, j)f(0,j) 以及 f(i, 0)f(i,0) 都设置为边界条件，它们的值均为 1。
func uniquePaths(m int, n int) int {
	dp := make([][]int,m)

	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}

	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m ; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
