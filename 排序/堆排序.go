package main

import "fmt"

func main() {
	arr := []int{2, 1, 123, 234, 56, 1, 324, 5, 6, 4, 543, 4, 5, 34, 5, 345, 34, 5}
	HeapSort(arr)
	fmt.Println(arr)
}

func HeapSort(arr []int) {
	arrLen := len(arr)
	// 建立大顶堆
	for i := arrLen/2 - 1; i >= 0; i-- {
		//从第一个非叶子结点从下至上，从右至左调整结构
		arrJustDown(arr, i, arrLen)
	}

	//2.调整堆结构+交换堆顶元素与末尾元素
	for j := arrLen - 1; j > 0; j-- {
		arr[0], arr[j] = arr[j], arr[0]
		arrJustDown(arr, 0, j)
	}
}

func arrJustDown(arr []int, i int, length int) {
	temp := arr[i]                              //先取出当前元素i
	for k := i*2 + 1; k < length; k = k*2 + 1 { //从i结点的左子结点开始，也就是2i+1处开始
		if k+1 < length && arr[k] < arr[k+1] { //如果左子结点小于右子结点，k指向右子结点
			k++
		}
		if arr[k] > temp { //如果子节点大于父节点，将子节点值赋给父节点（不用进行交换）
			arr[i] = arr[k]
			i = k
		} else {
			break
		}
	}
	arr[i] = temp //将temp值放到最终的位置
}
func arrJustDown2(arr []int, root, n int) {
	parent := root
	child := parent*2 + 1
	for child < n {
		if child+1 < n && arr[child+1] > arr[child] {
			child++
		}
		if arr[child] > arr[parent] {
			arr[child], arr[parent] = arr[parent], arr[child]
			parent = child
			child = parent*2 + 1
		} else {
			break
		}
	}
}
