package main

import (
	"fmt"
	"sort"
)

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

func main() {
	//回溯

	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {

	sort.Ints(nums)
	var res [][]int
	n := len(nums)
	used := make([]bool, n)



	var dfs func(index int, path []int)
	dfs = func(index int, path []int) {
		if len(path) == n {
			tmp := make([]int, n)
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := index; i < n; i++ {

			if !used[i] {
				used[i] = true
				path = append(path, nums[i])
				dfs(index, path)
				path = path[:len(path)-1]
				used[i] = false
			}

		}
	}
	dfs(0, []int{})
	return res
}
