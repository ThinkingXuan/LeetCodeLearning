package main

//739. 每日温度
//给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用 0 来代替。

func dailyTemperatures(temperatures []int) []int {
	// 求更高，递减栈
	res := make([]int, len(temperatures))
	var st []int

	for i := 0; i < len(temperatures); i++ {

		for len(st) > 0 && temperatures[st[len(st)-1]] < temperatures[i] {
			top := st[len(st)-1]
			res[top] = i - top
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return res
}
