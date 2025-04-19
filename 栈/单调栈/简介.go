package main

import "fmt"

/*
单调栈
定义：栈内元素单调按照递增(递减)顺序排列的栈。

单调递增栈：
①在一个队列中针对每一个元素从它右边寻找第一个比它小的元素
②在一个队列中针对每一个元素从它左边寻找第一个比它小的元素

单调递减栈：
①在一个队列中针对每一个元素从它右边寻找第一个比它大的元素
②在一个队列中针对每一个元素从它左边寻找第一个比它大的元素

单调栈何时用：为任意一个元素找左边和右边第一个比自己大/小的位置用单调栈.
由于每个元素最多各自进出栈一次,复杂度是O(n).
*/

//单调递减栈模板： 找到当前元素左边第一个小于它的的元素

func main() {
	var n int
	fmt.Scan(&n)
	var stack []int
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		for len(stack) > 0 && stack[len(stack)-1] > x {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			fmt.Println(-1)
		} else {
			fmt.Println(stack[len(stack)-1])
		}

		stack = append(stack, x)
	}
	return
}
