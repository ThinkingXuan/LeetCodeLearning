package main

// 二维转一维
func numTo(x, y, n int) int {
	return x*n + y
}

func numIslands(grid [][]byte) int {

	dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	m, n := len(grid), len(grid[0])

	// 重新初始化
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			father[numTo(i, j, m)] = i
		}
	}

	for i := 0; i <= m*n; i++ {
		father[i] = i
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if grid[i][j] == '0' {
				continue
			}
			for _, d := range dir {
				nx := i + d[0]
				ny := j + d[1]

				if nx < 0 || nx >= m || ny < 0 || ny >= n {
					continue
				} else if grid[nx][ny] == '1' {
					// 加入并查集
					join(numTo(i, j, n), numTo(nx, ny, n))
				}
			}
		}
	}

	res := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//如果字符为1，且根是自己本身，说明找到一个根，ans++
			if grid[i][j] == '1' && find(numTo(i, j, n)) == numTo(i, j, n) {
				res++
			}
		}
	}

	return res
}
