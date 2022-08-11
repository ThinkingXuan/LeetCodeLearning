package main

import (
	"fmt"
	"unsafe"
)

//比如在64位系统上，需要对一个[]float64切片进行高速排序，我们可以将它强制转为[]int整数切片，然后以整数的方式进行排序（因为float64遵循IEEE754浮点数标准特性，当浮点数有序时对应的整数也必然是有序的）。
func main() {
	a := []float64{2.5, 2.4, 2.3, 2.1, 2, 2.6}
	SortFloat64FastV1(a)
	fmt.Println(a)
}

func SortFloat64FastV1(a []float64) {
	// 强制类型转换
	var b = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]

	fmt.Println(b)

}
