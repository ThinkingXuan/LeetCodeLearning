package main

// mul 快速乘法  a乘数， k被乘数
// 重要
func mul(a int, k int) int {
	ans := 0
	for k > 0 {
		if (k & 1) == 1 {
			// 当前最低位为1，结果里加上a
			ans += a
		}
		// 被乘数右移1位，相当于除以2
		k >>= 1
		// 乘数倍增，相当于乘以2
		a += a
	}
	return ans
}

func mul2(a int, k int) int {
	ans := 0
	for ; k > 0; k >>= 1 {
		if k&1 == 1 {
			ans += a
		}
		a += a
	}
	return ans
}
