package main

import "fmt"

// 回溯题解
//https://leetcode-cn.com/problems/subsets/solution/shou-hua-tu-jie-zi-ji-hui-su-fa-xiang-jie-wei-yun-/

/*
22. 括号生成
代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]
*/

func main() {
	//fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis2(3))
}


/*方法一：推荐
选括号：

默认有n个左括号，n个右括号

先选左括号，一直能选，一直选，不能选时选右括号，保证右括号大于左括号
使用path来记录括号对

结束条件： 2 *n == len(Path)

*/
func generateParenthesis(n int) []string {

	var res []string
	var dfs func(lRemain, rRemain int, path string)

	dfs = func(lRemain, rRemain int, path string) {
		if 2*n == len(path) {
			res = append(res, path)
			return
		}

		if lRemain > 0 {
			dfs(lRemain-1, rRemain, path+"(")
		}
		if rRemain > lRemain {
			dfs(lRemain, rRemain-1, path+")")
		}
	}
	dfs(n, n, "")
	return res
}

/*
方法二: 每一次括号都有两种选择，找出所有的可能的情况 把每次的情况记为path，然后通过括号匹配算法来判断path是否正确，
*/

var parent = []string{"(", ")"}

func generateParenthesis2(n int) []string {
	var res []string
	var dfs func(n, index int, path string)

	dfs = func(n, index int, path string) {
		if index == n {
			if judge(path) {
				res = append(res, path)
			}
		}
		for i := 0; i < 2; i++ {
			dfs(n, index+1, path+parent[i])
		}
	}
	dfs(2*n, 0, "")
	return res
}

// 括号比配
func judge(s string) bool {

	n := len(s)

	if n <=1 && s[0] == ')' {
		return false
	}
	var stack []byte
	flag := true

	for i:=0; i<len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, s[i])
		}else {
			if len(stack) == 0 {
				flag = false
				break
			}
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if left == '(' {
				flag = false
			}
		}
	}

	if flag {
		if len(stack) != 0 {
			flag = false
		}
	}
	return flag
}
