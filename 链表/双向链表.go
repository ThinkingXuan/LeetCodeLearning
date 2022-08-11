package main

import (
	"container/list"
	"fmt"
)

func printList(ll *list.List) {
	fmt.Println("================================")
	for e := ll.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println("\n================================")
}

func main() {

	// 初始化
	ll := list.New()

	// 添加元素
	for i := 4; i >= 0; i-- {
		ll.PushFront(i)
	}
	for i := 5; i < 10; i++ {
		ll.PushBack(i)
	}
	printList(ll)

	var ele *list.Element //找到元素值为5的元素
	for e := ll.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
		if e.Value.(int) == 5 {
			ele = e
			break
		}
	}
	// 插入 InsertBefore
	ll.InsertBefore(4.9, ele)

	// 插入 InsertAfter
	ll.InsertAfter(5.1, ele)

	printList(ll)

	// 遍历方法 Front() Back()
	fmt.Println(ll.Front().Value) // 获取头元素
	fmt.Println(ll.Back().Value)  // 获取尾元素

	// Next() Prev()
	fmt.Println(ele.Next().Value)
	fmt.Println(ele.Prev().Value)

	// Remove,删除尾元素
	ll.Remove(ll.Back())
	printList(ll)

	//移动元素 MoveToFront MoveToBack
	// 将value = 5的元素移动到头部（尾部），
	ll.MoveToFront(ele)
	//ll.MoveToBack(ele)
	printList(ll)

	// 移动元素指定的元素前后  MoveBefore MoveAfter

	// 获取长度
	fmt.Println(ll.Len())
}
