package main

import (
	"container/heap"
	"fmt"
	"testing"
)

// 完全重写所有方法
// 来自官方heap测试代码
type hp []int

func (h *hp) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *hp) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *hp) Len() int {
	return len(*h)
}

func (h *hp) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func TestInit(t *testing.T) {

	h := &hp{}

	// 倒叙填充值
	for i := 5; i > 0; i-- {
		*h = append(*h, i)
	}

	// init
	heap.Init(h)
	fmt.Println(h)

	// push
	heap.Push(h, 2) // 会自动调整
	fmt.Println(h)

	// pop
	v := heap.Pop(h)
	fmt.Printf("v = %v h = %v \n", v, h)

	// fix 作用是手动调整，比如改变hp里面下标i的值，然后使用它的进行调整
	(*h)[0] = 8
	heap.Fix(h, 2)
	fmt.Println(h)

	// remove 手动删除特定位置的元素
	i := 2
	heap.Remove(h, i)
	fmt.Println(h)

}
