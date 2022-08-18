package main

import "fmt"

//https://blog.csdn.net/qq_45914558/article/details/107385862

/*输入一个长度为n的整数序列。

接下来再输入m个询问，每个询问输入一对l, r。

对于每个询问，输出原序列中从第l个数到第r个数的和。

输入格式
第一行包含两个整数n和m。

第二行包含n个整数，表示整数数列。

接下来m行，每行包含两个整数l和r，表示一个询问的区间范围。

输出格式
共m行，每行输出一个询问的结果。*/

//1 2 3 4 5 6 7 8 9 10
//0 1 3 6 10 15
//0 1 2 3 4
// [3, 4] 5 - 3

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	nums := make([]int, n+1)
	preSum := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&nums[i])
		preSum[i] = preSum[i-1] + nums[i]
	}

	for i := 0; i < m; i++ {
		l, r := 0, 0 // 从1开始,第l和r个, [l, r]
		fmt.Scan(&l, &r)
		fmt.Println(preSum[r] - preSum[l-1]) // 下标从0开始, preSum[r+1] - preSum[l]
	}
}
