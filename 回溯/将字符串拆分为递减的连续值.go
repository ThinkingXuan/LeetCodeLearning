package main

import (
	"fmt"
)

/*

给你一个仅由数字组成的字符串 s 。

请你判断能否将 s 拆分成两个或者多个 非空子字符串 ，使子字符串的 数值 按 降序 排列，且每两个 相邻子字符串 的数值之 差 等于 1 。

例如，字符串 s = "0090089" 可以拆分成 ["0090", "089"] ，数值为 [90,89] 。这些数值满足按降序排列，且相邻值相差 1 ，这种拆分方法可行。
另一个例子中，字符串 s = "001" 可以拆分成 ["0", "01"]、["00", "1"] 或 ["0", "0", "1"] 。然而，所有这些拆分方法都不可行，因为对应数值分别是 [0,1]、[0,1] 和 [0,0,1] ，都不满足按降序排列的要求。
如果可以按要求拆分 s ，返回 true ；否则，返回 false 。

子字符串 是字符串中的一个连续字符序列。
*/

func main() {
	s := "0090089"
	fmt.Println(splitString(s))
}

/*
	思想： 因为字符串之间有0影响，比如0090 089，他们之间不能通过简单的比较就可以判断出来，需要先去除0，然后在比较，较为麻烦。
	因为数组直接直接计算数值  “0090” => 90 "089"=> 89,可以直接判断 他们之间相差 1，所以最好转换为数组去做。

	下面就是DFS搜索，可以画一个搜索搜索树
 						      []
			0 		0 		9 		0 		 0 	 	 8			9

	搜索位置从0开始，找到一个数时，然后从他本身的下一个数进行搜索dfs(i+1,....)
	搜索过程中，需要计算搜索过程中字符串的整数值  cur = 10 *cur + int(num[i) - '0'
	并且需要知道他上一个值的大小，如果比上一个值小1（用户剪枝），然后继续搜索
	还需要一个count来计算找到了几个。默认为0个，等搜索完成后，如果大于1个，就搜索成功。
*/
func splitString(s string) bool {
	if len(s) == 0 {
		return false
	}

	arr := []byte(s)

	var dfs func(index, count int, pre int) bool

	dfs = func(index, count int, pre int) bool {
		if index == len(arr) {
			return count > 1
		}
		var cur = 0
		for i := index; i < len(arr); i++ {
			cur = 10*cur + int(arr[i]) - '0'
			if count == 0 || cur == pre-1 {
				if dfs(i+1, count+1, cur) {
					return true
				}
			}
		}
		return false
	}

	return dfs(0, 0, 0)
}
