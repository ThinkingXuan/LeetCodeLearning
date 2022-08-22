package main

import (
	"fmt"
	"sort"
)

// sort.Search()
// sort.SearchInts()
// sort.SearchStrings()
// sort.SearchFloat64s

func main() {
	// 先有序
	//func Search(n int, f func(int) bool) int
	//采用二分法搜索找到[0, n)区间内最小的满足f(i)==true的值i
	//未找到时的返回值不是-1，而是n。这一点和strings.Index等函数不同。
	//并返回能够使f(i)=true的最小的i

	//https://leetcode-cn.com/problems/B1IidL/
	fmt.Println(peakIndexInMountainArray([]int{1, 3, 5, 4, 2})) // 找峰顶下标

	//sort.SearchInts(arr,x int)
	// 找a[i] >= x的最小下标

	//sort.SearchStrings(n int x string)
	//sort.SearchFloat64s(n int x float64)
	// 同理

	// sort.SearchInts 的使用技巧

	nums := []int{}
	target := 10

	lowerBound := sort.SearchInts(nums, target)

	upbound := sort.Search(len(nums), func(i int) bool {
		return nums[i] > target
	})

	// 第二种写法
	// upbound := sort.SearchInts(nums, target+1)

	_ = []interface{}{lowerBound, upbound - 1}
}

func peakIndexInMountainArray(arr []int) int {
	// 因为结果不可能是尾部
	// sort.Search(n, func(i int)bool)  满足f(i)== true的最小的i，没有搜索到返回n
	return sort.Search(len(arr)-1, func(i int) bool { return arr[i] > arr[i+1] })
}
