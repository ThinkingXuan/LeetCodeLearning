package main

//https://leetcode-cn.com/problems/number-of-islands/
//https://leetcode-cn.com/problems/number-of-islands/solution/number-of-islands-shen-du-you-xian-bian-li-dfs-or-/

// dfs
func numIslands(grid [][]byte) int {
	nr, nc := len(grid), len(grid[0])
	res := 0

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= nr || c < 0 || c >= nc || grid[r][c] == '0' {
			return
		}

		grid[r][c] = '0'

		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			if grid[i][j] == '1' {
				res++
				dfs(i, j)
			}
		}
	}

	return res

}

// bfs
func numIslands2(grid [][]byte) int {
	nr, nc := len(grid), len(grid[0])
	res := 0

	var bfs func(r, c int)
	bfs = func(r, c int) {
		var queue [][]int
		queue = append(queue, []int{r, c})

		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			r, c := cur[0], cur[1]

			if r >= 0 && r < nr && c >= 0 && c < nc && grid[r][c] == '1' {
				grid[r][c] = '0'
				queue = append(queue, []int{r - 1, c})
				queue = append(queue, []int{r + 1, c})
				queue = append(queue, []int{r, c - 1})
				queue = append(queue, []int{r, c + 1})
			}
		}
	}
	for i := 0; i < nr; i++ {
		for j := 0; j < nc; j++ {
			if grid[i][j] == '1' {
				res++
				bfs(i, j)
			}
		}
	}
	return res
}
