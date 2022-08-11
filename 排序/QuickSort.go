package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	arr := []int{2, 1, 123, 234, 56, 1, 324, 5, 6, 4, 543, 4, 5, 34, 5, 345, 34, 5}

	quickSort1(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// 需要记住的
func quickSort1(arr []int, L, R int) {
	left, right := L, R
	if left >= right {
		return
	}

	// 随机选择轴
	//randVal := left + rand.Intn(right-left+1)
	//arr[left], arr[randVal] = arr[randVal], arr[left]

	pivot := arr[left]
	for left < right {
		for left < right && arr[right] >= pivot {
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] <= pivot {
			left++
		}
		arr[right] = arr[left]
	}

	arr[left] = pivot
	quickSort1(arr, L, left-1)
	quickSort1(arr, left+1, R)
}

// QuickSort4 标准快速排序
func QuickSort4(nums []int, low, high int, k int) int {

	partition := randomPartition(nums, low, high)

	if partition == k {
		return nums[k-1]
	} else if partition < k {
		return QuickSort4(nums, partition+1, high, k)
	}

	return QuickSort4(nums, low, partition-1, k)
}
func randomPartition(nums []int, low, high int) int {
	rand.Seed(time.Now().UnixNano())
	picked := rand.Int() % (high - low + 1)
	nums[picked], nums[low] = nums[picked], nums[low]
	return partition(nums, low, high)
}

func partition(nums []int, low, high int) int {
	povit := nums[low]
	index := low
	for i := low + 1; i <= high; i++ {
		if nums[i] >= povit {
			nums[i], nums[index+1] = nums[index+1], nums[i]
			index++
		}
	}
	nums[low], nums[index] = nums[index], nums[low]
	return index
}

// 官方方法
func quickSort2(arr []int, low, high int) {

	if low > high {
		return
	}
	pivot := arr[low]

	index := low

	for i := index + 1; i <= high; i++ {
		if arr[i] <= pivot {
			arr[index+1], arr[i] = arr[i], arr[index+1]
			index++
		}
	}

	arr[low], arr[index] = arr[index], arr[low]
	quickSort2(arr, low, index-1)
	quickSort2(arr, index+1, high)
}

// go语言特有方法
func quickSort(s []int) []int {
	if len(s) < 2 {
		return s
	}
	v := s[0]
	var left, right []int
	for _, e := range s[1:] {
		if e <= v {
			left = append(left, e)
		} else {
			right = append(right, e)
		}
	}
	return append(append(quickSort(left), v), quickSort(right)...)
}

//   常规方法
func QuickSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	part := partition2(arr, low, high)
	QuickSort(arr, low, part-1)
	QuickSort(arr, part+1, high)
}

func partition2(arr []int, low, high int) int {

	povit := arr[low]

	i := low + 1
	j := high

	for true {
		for arr[i] < povit {
			i++
			if i == high {
				break
			}
		}

		for arr[j] > povit {
			j--
			if j == low {
				break
			}
		}

		if i >= j {
			break
		}

		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[low], arr[j] = arr[j], arr[low]

	return j
}

func quickSort3(arr []int, low, high int) {
	if low > high {
		return
	}

	pivod := arr[low]
	index := low

	for i := index + 1; i <= high; i++ {
		if arr[i] < pivod {
			arr[index+1], arr[i] = arr[i], arr[index+1]
			index++
		}
	}
	arr[index], arr[low] = arr[low], arr[index]

	quickSort3(arr, low, index-1)
	quickSort3(arr, index+1, high)
}

func findKthLargest(nums []int, k int) int {

	var ans *int
	quicksort(nums, 0, len(nums)-1, k, ans)
	return *ans
}
func quicksort(nums []int, low, high int, k int, ans *int) {
	rand.Seed(time.Now().UnixNano())
	picked := rand.Int()%(high-low+1) + low
	nums[picked], nums[low] = nums[low], nums[picked]

	povit := nums[low]
	index := low

	for i := low + 1; i <= high; i++ {
		if nums[i] >= povit {
			nums[index+1], nums[i] = nums[i], nums[index+1]
			index++
		}
		nums[index], nums[low] = nums[low], nums[index]

		if k < index {
			quicksort(nums, low, index-1, k, ans)
		} else if k == index {
			ans = &nums[index]
			return
		} else {
			quicksort(nums, index+1, high, k-index+1, ans)
		}

	}
}
