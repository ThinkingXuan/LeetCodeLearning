package main

import "testing"

// https://leetcode-cn.com/problems/island-perimeter/
// 对于一个陆地格子的每条边，它被算作岛屿的周长当且仅当这条边为网格的边界或者相邻的另一个格子为水域。 因此，我们可以遍历每个陆地格子，看其四个方向是否为边界或者水域，
// 如果是，将这条边的贡献（即 11）
// 463

// dfs
func islandPerimeter(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == 0 {
			res++
			return
		}

		// 陆地里面遍历过的的不要计算。。
		if grid[r][c] == 2 {
			return
		}

		grid[r][c] = 2
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
			}
		}
	}
	return res
}

// bfs
func islandPerimeter2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0

	var bfs func(r, c int)
	bfs = func(r, c int) {
		queue := make([][]int, 0)
		queue = append(queue, []int{r, c})

		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]

			x, y := cur[0], cur[1]
			if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 {
				res++
			} else if grid[x][y] == 2 {
				continue
			} else {
				grid[x][y] = 2
				queue = append(queue, []int{x + 1, y})
				queue = append(queue, []int{x - 1, y})
				queue = append(queue, []int{x, y + 1})
				queue = append(queue, []int{x, y - 1})
			}
		}

	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				bfs(i, j)
			}
		}
	}

	return res

}

func TestIslandPerimeter(t *testing.T) {
	testCase := []struct {
		in   [][]int
		want int
	}{
		{[][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}, 16},
		{[][]int{{4}}, 4},
		{[][]int{{1, 0}}, 4},
	}

	for _, tc := range testCase {
		res := islandPerimeter(tc.in)
		if res != tc.want {
			t.Errorf("checkInclusion: %v, want %v", res, tc.want)
		}
	}
}
