package main

import (
	"math"
)

func main() {
	sieveMu2()
}

// 寻找num之内的所有素数, 直接试除法
func meth01(num int) []int {
	res := []int{1}

	for i := 2; i < num; i++ {
		j := 2
		for ; j < i; j++ {
			if i%j == 0 {
				break
			}
		}

		if i == j {
			res = append(res, i)
		}
	}
	return res
}

//寻找num之内的所有素数, 试除法， 使用Sqrt优化
func meth02(num int) []int {

	res := []int{1}

	for i := 2; i < num; i++ {
		j := 2
		for ; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				break
			}
		}

		if j > int(math.Sqrt(float64(i))) {
			res = append(res, i)
		}
	}
	return res
}

//寻找num之内的所有素数, 试除法， 使用Sqrt优化
func meth03(num int) []int {
	res := []int{1, 2}
	for i := 3; i < num; i = i + 2 {
		j := 2
		for ; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				break
			}
		}

		if j > int(math.Sqrt(float64(i))) {
			res = append(res, i)
		}
	}
	return res
}

// 线性筛
func sieveMu() []int {
	const mx int = 1e6
	mu := [mx + 1]int{1: 1} // int8
	primes := []int{}
	vis := [mx + 1]bool{}
	for i := 2; i <= mx; i++ {
		if !vis[i] {
			mu[i] = -1
			primes = append(primes, i)
		}
		for _, p := range primes {
			v := p * i
			if v > mx {
				break
			}
			vis[v] = true
			if i%p == 0 {
				mu[v] = 0
				break
			}
			mu[v] = -mu[i]
		}
	}
	return mu[:]
}

// 埃氏筛写法 (效率更高)  值为-1为质数，其他不是质数
func sieveMu2() []int {
	const mx int = 15
	//0 1 -1 -1 0 -1 1 -1 0 0 1 -1 0 -1 1 1
	mu := [mx + 1]int{1: 1} // int8
	for i := 1; i <= mx; i++ {
		for j := i * 2; j <= mx; j += i {
			mu[j] -= mu[i]
		}
	}
	return mu[:]
}
