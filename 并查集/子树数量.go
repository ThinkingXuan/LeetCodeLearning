package main

//https://leetcode-cn.com/problems/number-of-provinces/

func findCircleNum(isConnected [][]int) int {

	initialize()
	for i := 0; i < len(isConnected); i++ {
		// 因为是无向图，所以只传一半得节点
		for j := i + 1; j < len(isConnected[i]); j++ {
			if isConnected[i][j] == 1 {
				join(i, j)
			}
		}
	}
	// 统计连通分量
	res := connectSize(len(isConnected))

	return res
}
