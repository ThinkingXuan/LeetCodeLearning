package main

import (
	"container/heap"
	"fmt"
	"math"
)

//给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案

//桶排序
func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
func topKFrequent(nums []int, k int) []int {

	maxn := math.MinInt64
	// 将数放于map中，统计每一个数的个数,并统计重复数最大的数量
	mp := make(map[int]int, k)
	for _, v := range nums {
		mp[v]++
		if mp[v] > maxn {
			maxn = mp[v]
		}
	}

	// 建立桶
	hastop := make([][]int, maxn+1)
	for key, val := range mp {
		hastop[val] = append(hastop[val], key)
	}

	res := make([]int, 0)
	// 排序桶，并得出结果
	for i := maxn; i >= 0; i-- {
		res = append(res, hastop[i]...)
		k -= len(hastop[i])
		if k == 0 {
			break
		}
	}

	return res
}

//堆排序

func topKFrequent2(nums []int, k int) []int {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}

	h := &IHeap{}
	heap.Init(h)

	for key, value := range occurrences {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//快速排序
