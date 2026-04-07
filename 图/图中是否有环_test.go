package main

import "testing"

// 判断邻接矩阵代表的图是否有环。
// 使用 DFS 深度优先遍历。
// 如果访问到了一个已经访问过的邻居 且不是父节点，则说明存在环
func hasCycleDFS(graph [][]int) bool {
	n := len(graph)

	visited := make([]bool, n)

	var dfs func(node, parent int) bool
	dfs = func(node, parent int) bool {
		visited[node] = true
		for v := 0; v < n; v++ {
			if graph[node][v] == 1 && !visited[node] {

				if !visited[node] {
					if dfs(v, node) {
						return true
					}
				} else if v != parent {
					// 如果已经访问，并且不是父节点，说明有环
					// 如果 v 被访问过，且不是我从它来的“父节点”
					// 那说明从另一条路径绕回来了 => 构成环！
					return true
				}
			}
		}

		return false
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			if dfs(i, -1) {
				return true
			}
		}
	}

	return false
}

// 有向图使用DFS+三色标记法，判断是否有环
func HasCycle(graph [][]int) bool {
	const (
		WHITE = 0
		GREY  = 1
		BLACK = 2
	)

	color := make([]int, len(graph))

	var dfs func(node int) bool
	dfs = func(node int) bool {
		color[node] = GREY
		for neighbor := 0; neighbor < len(graph); neighbor++ {
			if graph[node][neighbor] == 0 {
				continue
			}
			if color[neighbor] == GREY {
				// 遇到灰色节点，有环
				return true
			}
			if color[neighbor] == WHITE {
				if dfs(neighbor) {
					return true
				}
			}
		}
		color[node] = BLACK // 访问完成
		return false
	}

	for i := 0; i < len(graph); i++ {
		if color[i] == WHITE {
			if dfs(i) {
				return true
			}
		}
	}
	return false
}

func TestHashCircle(t *testing.T) {
	testCase := []struct {
		in   [][]int
		want bool
	}{
		{
			in: [][]int{
				{0, 1, 0, 1},
				{1, 0, 1, 0},
				{0, 1, 0, 1},
				{1, 0, 1, 0}},
			want: true,
		},
		{
			in: [][]int{
				{0, 1, 0, 0},
				{1, 0, 1, 0},
				{0, 1, 0, 1},
				{0, 0, 1, 0}},
			want: false,
		},
	}

	for _, c := range testCase {
		if hasCycleDFS(c.in) != c.want {
			t.Log("fail", c.want, c.in)
		}
	}
}
