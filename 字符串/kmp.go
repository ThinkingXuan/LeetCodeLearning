package main

import (
	"fmt"
)

// 求Next数组
func getNext(s string) []int {
	next := make([]int, len(s))
	j := -1 // j从-1开始
	next[0] = j

	for i := 1; i < len(s); i++ { // i从1开始
		for j >= 0 && s[i] != s[j+1] { // 前缀不相同
			j = next[j] // 向前回退
		}

		if s[i] == s[j+1] { // 找到相同的前后缀
			j++
		}
		next[i] = j // 将j（前缀的长度）赋给next[i]
	}
	return next
}

func main() {
	fmt.Println(getNext("aabaaf"))
}
