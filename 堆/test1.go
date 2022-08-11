package main

import (
	"container/heap"
	"sort"
)

// 使用sort.IntSlice重写部分方法
// go语言大根堆实现
type myHeap struct {
	sort.IntSlice
}

// Less 因为默认是小根堆，所以要重新下面方法
func (h myHeap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *myHeap) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *myHeap) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func main() {
	sum := []int{6, 2, 4, 6, 8}
	h := &myHeap{sum}
	heap.Init(h)

	k := 3
	for ; k > 0; k-- {
		h.IntSlice[0] -= h.IntSlice[0] / 2
		heap.Fix(h, 0)

	}

	//for _,v := range h.IntSlice {
	//
	//}
}

func printPrime(prefix string) {

}
