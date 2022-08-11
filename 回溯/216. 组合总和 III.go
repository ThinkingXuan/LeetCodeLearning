package main

import (
	"fmt"
)

//找出所有相加之和为n的k个数的组合。组合中只允许含有 1 -9 的正整数，并且每种组合中不存在重复的数字。
//
//说明：
//
//所有数字都是正整数。
//解集不能包含重复的组合
//

func main() {
	fmt.Println(combinationSum3(3, 10))
}

func combinationSum3(k int, n int) [][]int {

	ans := [][]int{}
	var dfs func(start int, path []int, sum int)

	dfs = func(start int, path []int, sum int) {
		if k == len(path) {
			if n == sum {
				tmp := make([]int, len(path))
				copy(tmp, path)
				ans = append(ans, tmp)
			}
		}
		for i := start; i <= 9; i++ {
			path = append(path, i)
			dfs(i+1, path, sum+i)
			path = path[:len(path)-1]
		}
	}

	dfs(1, []int{}, 0)
	return ans
}
