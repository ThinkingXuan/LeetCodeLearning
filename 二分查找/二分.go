package main

import "fmt"

//https://www.cnblogs.com/kyoner/p/11080078.html

// 正常的二分查找（返回索引）
func test1(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2
		// mid := (left + right) >> 1

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
	//nums[left] == target ? left : -1;
}

// 搜索左侧边界
// 等于 right = mid
// 大于 right = mid
// 小于 left = mid+1
// 返回 left
func leftBound(nums []int, target int) int {
	left := 0
	right := len(nums)

	for left < right {
		mid := left + (right-left)/2
		// mid := (left + right) >> 1
		if nums[mid] == target {
			right = mid
		} else if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return left
	//nums[left-1] == target ? (left-1) : -1;
}

// 搜索右侧边界
// 等于 left = mid+1
// 大于 right = mid
// 小于 left = mid+1
// 返回 left-1
func rightBound(nums []int, target int) int {
	left := 0
	right := len(nums)

	for left < right {
		mid := left + (right-left)/2
		// mid := (left + right) >> 1
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return left - 1
}

func main() {
	nums := []int{1, 3, 6, 7, 9, 12, 13}
	fmt.Println(rightBound(nums, 1) + 1)
	fmt.Println(rightBound(nums, 8) + 1)

	fmt.Println(rightBound(nums, 11) + 1)
	fmt.Println(rightBound(nums, 14) + 1)

}

//func main() {
//	//nums1 := []int{1, 2, 4, 5, 6, 7, 8}
//	//fmt.Println(test1(nums1, 2))
//
//	// 使用左右边界写法
//	nums2 := []int{1, 2, 5, 10, 16, 18, 20}
//	index := test2(nums2, 3)
//	fmt.Println(nums2[index])
//	index = test3(nums2, 3)
//	fmt.Println(nums2[index])
//
//	//SeachInts
//	fmt.Println(nums2[sort.SearchInts(nums2, 3)-1])
//	fmt.Println(nums2[sort.SearchInts(nums2, 3+1)])
//
//}
