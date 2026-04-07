package main

import (
	"fmt"
)

// 图的表示方法
// 邻接表

/*
// 邻接矩阵
对于无向图：

如果顶点 i 和 j 之间有边相连，则 matrix[i][j] = 1 且 matrix[j][i] = 1
如果没有边相连，则为 0，对角线元素（matrix[i][i]）通常设为 0（除非允许自环）
邻接矩阵:
顶点: 0, 1, 2, 3
边: 0-1, 1-2, 2-3, 3-0
[
	[0, 1, 0, 1],
	[1, 0, 1, 0],
	[0, 1, 0, 1],
	[1, 0, 1, 0]
]
*/

// 把边转化为邻接矩阵(无向图)
func getMatrix(nums [][2]int) [][]int {
	n := len(nums)

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for _, v := range nums {
		matrix[v[0]][v[1]] = 1
		matrix[v[1]][v[0]] = 1
	}

	return matrix
}

// 把边转化为邻接表(无向图)
func getTable(nums [][2]int) [][]int {
	adj := make([][]int, len(nums))
	for i := range adj {
		adj[i] = []int{}
	}

	for _, v := range nums {
		adj[v[0]] = append(adj[v[0]], v[1])
		adj[v[1]] = append(adj[v[1]], v[0])
	}
	return adj
}

func main() {

	input := [][2]int{
		{0, 1}, {1, 2}, {2, 3}, {3, 0},
	}

	fmt.Println("------------------邻接矩阵-------------------")
	matrix := getMatrix(input)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}

	dfsRes := getNodeListForDFS(matrix)
	fmt.Println("DFS: ", dfsRes)

	bfsRes := getNodeListForBFS(matrix)
	fmt.Println("BFS: ", bfsRes)

	fmt.Println("------------------邻接表-------------------")

	adj := getTable(input)

	for i := 0; i < len(adj); i++ {
		fmt.Printf("node:%d => ", i)
		for j := 0; j < len(adj[i]); j++ {
			fmt.Print(adj[i][j], " ")
		}
		fmt.Println()
	}

	dfsRes = getAdjNodeListForDFS(adj)
	fmt.Println("DFS: ", dfsRes)

	bfsRes = getAdjNodeListForBFS(adj)
	fmt.Println("BFS: ", bfsRes)
}

// 遍历所有节点
func getNodeListForDFS(graph [][]int) []int {
	// dfs
	visited := make([]bool, len(graph))
	var res []int

	var dfs func(start int)

	dfs = func(start int) {

		res = append(res, start)

		visited[start] = true
		for i := 0; i < len(graph); i++ {
			if graph[start][i] == 1 && !visited[i] {
				dfs(i)
			}
		}
	}
	dfs(0)

	return res
}

// 遍历所有节点
func getNodeListForBFS(graph [][]int) []int {
	// bfs
	var res []int
	visited := make([]bool, len(graph))

	var queue []int
	// 0节点入队
	queue = append(queue, 0)
	visited[0] = true

	for len(queue) > 0 {
		// 入队
		cur := queue[0]
		queue = queue[1:]
		res = append(res, cur)

		// 所有邻接点入队
		for i := 0; i < len(graph); i++ {
			if graph[cur][i] == 1 && !visited[i] {
				visited[i] = true
				queue = append(queue, i)
			}
		}
	}

	return res
}

func getAdjNodeListForDFS(adj [][]int) []int {
	visited := make([]bool, len(adj))
	var res []int

	var dfs func(start int)

	dfs = func(start int) {

		res = append(res, start)

		visited[start] = true
		for _, v := range adj[start] {
			if !visited[v] {
				dfs(v)
			}
		}
	}
	dfs(0)

	return res
}

// 遍历所有节点
func getAdjNodeListForBFS(adj [][]int) []int {
	// bfs
	var res []int
	visited := make([]bool, len(adj))

	var queue []int
	// 0节点入队
	queue = append(queue, 0)
	visited[0] = true

	for len(queue) > 0 {
		// 出队
		cur := queue[0]
		queue = queue[1:]
		res = append(res, cur)

		// 所有邻接点入队
		for _, v := range adj[cur] {
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}

	return res
}
