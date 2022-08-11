package main

import (
	"fmt"
)

// quickPow 快速幂
// a的n次幂
//
//  (2,10)     (2,5)     (2,2) (2,1) (2,0)
//   32*32        4*4*2=32     4      2      1
func quickPow(a float64, n int) float64 {
	if n == 0 {
		return 1
	}

	y := quickPow(a, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * a
}

//https://leetcode-cn.com/problems/powx-n/solution/50-powx-n-kuai-su-mi-qing-xi-tu-jie-by-jyd/
func quickPow2(a float64, n int) float64 {
	ans := 1.0
	for n > 0 {
		if n&1 == 1 { // n % 2 == 1
			ans = ans * a
		}
		a *= a
		n >>= 1 // n /= 2
	}
	return ans
}

func quickPow3(a int, n int) int {
	ans := 1
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			ans = ans * a
		}
		a *= a
	}
	return ans
}

func main() {
	fmt.Println(quickPow2(2, 10))

	var a []int
	n, _ := fmt.Scan(&a)
	fmt.Println(n, a)

}
