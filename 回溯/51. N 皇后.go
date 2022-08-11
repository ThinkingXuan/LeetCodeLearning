package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(solveNQueens(4))
}

func solveNQueens(n int) [][]string {
	bd := make([][]string, n)
	for i := range bd {
		bd[i] = make([]string, n)

		for j := range bd[i] {
			bd[i][j] = "."
		}
	}
	 ans := [][]string{}
	fmt.Printf("start: %p\n", ans)
	dfs(0, bd, ans, n)
	return ans
}

func dfs(r int, bd [][]string, ans [][]string, n int) {

	fmt.Printf("%p \n", ans)
	if r == n {
		temp := make([]string, n)
		for i := 0; i < n; i++ {
			temp[i] = strings.Join(bd[i], "")
		}
		// append之后就不是原来的ans的，内存地址变化了。
		ans = append(ans, temp)
		fmt.Printf("%p \n",ans)
		return
	}

	for c := 0; c < n; c++ {
		if isValid(r, c, n, bd) {
			bd[r][c] = "Q"
			dfs(r+1, bd, ans, n)
			bd[r][c] = "."
		}
	}
}

func isValid(r, c int, n int, bd [][]string) bool {
	for i := 0; i < r; i++ {
		for j := 0; j < n; j++ {
			if bd[i][j] == "Q" && (j == c || i+j == r+c || i-j == r-c) {
				return false
			}
		}
	}
	return true
}
