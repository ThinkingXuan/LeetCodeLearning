//在给定的二维二进制数组 A 中，存在两座岛。（岛是由四面相连的 1 形成的一个最大组。）
//
// 现在，我们可以将 0 变为 1，以使两座岛连接起来，变成一座岛。
//
// 返回必须翻转的 0 的最小数目。（可以保证答案至少是 1 。）
//
//
//
// 示例 1：
//
//
//输入：A = [[0,1],[1,0]]
//输出：1
//
//
// 示例 2：
//
//
//输入：A = [[0,1,0],[0,0,0],[0,0,1]]
//输出：2
//
//
// 示例 3：
//
//
//输入：A = [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
//输出：1
//
//
//
// 提示：
//
//
// 2 <= A.length == A[0].length <= 100
// A[i][j] == 0 或 A[i][j] == 1
//
//
// Related Topics 深度优先搜索 广度优先搜索 数组 矩阵 👍 271 👎 0

//1,1,1,1,1
//1,0,0,0,1
//1,0,1,0,1
//1,0,0,0,1
//1,1,1,1,1

package main

//leetcode submit region begin(Prohibit modification and deletion)
// 思路 dfs 找到一个岛，赋值为2，，同时加入队列。。然后BFS，找最短。
func shortestBridge(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	var queue [][]int

	var dfs func(r, c int)
	var bfs func() int

	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == 0 || grid[r][c] == 2 {
			return
		}
		queue = append(queue, []int{r, c})
		grid[r][c] = 2

		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	bfs = func() int {
		dir := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		res := 0

		for len(queue) > 0 {
			size := len(queue)
			res++

			for i := 0; i < size; i++ {
				cur := queue[0]
				queue = queue[1:]
				x, y := cur[0], cur[1]

				for _, d := range dir {
					newX, newY := x+d[0], y+d[1]
					if newX >= 0 && newX < m && newY >= 0 && newY < n && grid[newX][newY] != 2 {
						if grid[newX][newY] == 0 {
							queue = append(queue, []int{newX, newY})
							grid[newX][newY] = 2
						} else if grid[newX][newY] == 1 {
							return res - 1
						}
					}
				}
			}
		}
		return res
	}

s:
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				break s
			}
		}
	}

	return bfs()
}

//leetcode submit region end(Prohibit modification and deletion)
