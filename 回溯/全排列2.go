package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique([]int{1,1,2}))
}
func permuteUnique(nums []int) [][]int {


	sort.Ints(nums)

	n := len(nums)
	var res [][]int
	path := []int{}

	visit := make([]bool, n)

	var dfs func(index int)

	dfs = func(index int) {
		if len(path) == n {
			tmp := make([]int, n)
			copy(tmp ,path)
			res = append(res, tmp)
			return
		}

		for i := index; i < n; i++ {
			if visit[i] {
				continue
			}

			// 同一层，相同的值值访问一遍，
			//还行保证他前一个没有被访问过, 因为会进行回溯，访问到下一个相同的数值时，前一个visit已经被置为false
			if i-1>=0 && nums[i] == nums[i-1] && !visit[i-1] {
				continue
			}

			visit[i] = true
			path = append(path, nums[i])
			dfs(index)
			path = path[:len(path)-1]
			visit[i] = false
		}
	}
	dfs(0)
	return res
}
