package main

import "fmt"

// 归并排序（分治法）

func sortM(nums []int, left, right int, tmp []int) {
	if left >= right {
		return
	}

	mid := (left + right) / 2
	sortM(nums, left, mid, tmp)
	sortM(nums, mid+1, right, tmp)
	merge(nums, left, mid, right, tmp)

}

func merge(nums []int, left, mid, right int, tmp []int) {
	i := left    //左序列
	j := mid + 1 //右序列
	index := 0   //临时数组

	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp[index] = nums[i]
			i++
		} else {
			tmp[index] = nums[j]
			j++
		}
		index++
	}

	//将左边剩余元素填充进temp中
	for i <= mid {
		tmp[index] = nums[i]
		index++
		i++
	}

	//将右序列剩余元素填充进temp中
	for j <= right {
		tmp[index] = nums[j]
		index++
		j++
	}

	//将temp中的元素全部拷贝到原数组中
	index = 0
	for left <= right {
		nums[left] = tmp[index]
		left++
		index++
	}
}
func main() {
	nums := []int{7, 10, 1, 5, 62, 2, 6, 23, 4, 23, 7, 3}
	tmp := make([]int, len(nums))
	sortM(nums, 0, len(nums)-1, tmp)
	fmt.Println(nums)
}
