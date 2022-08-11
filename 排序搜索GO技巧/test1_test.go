package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{1, 2, 45, 223, 23, 2423, 43, 3}
	b := sort.IntsAreSorted(arr)

	fmt.Println(b)
}

func TestStringSort(t *testing.T) {
	str := []string{"12", "1111", "1234"}
	// 按照字典序
	sort.Strings(str)
	fmt.Println(str)
}

func TestSortSearch(t *testing.T) {
	//sort.Search 二分查找方式，必须升序排序， 查找失败后，会返回切片得长度
	arr := []int{1,2,3,4,5,6,7,8}
	_ = sort.Search(len(arr), func(i int) bool {
		return arr[i] <= 5
	})
	d1 := sort.SearchInts(arr,100)
	fmt.Println(d1)
}





