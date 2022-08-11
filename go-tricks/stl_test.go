package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var appendA = []int{1, 2, 3, 4, 5}

// 添加单个值 pushBack
func TestPushBack(t *testing.T) {
	appendA = append(appendA, 10)
	assert.Equal(t, appendA, []int{1, 2, 3, 4, 5, 10})
}

// 添加vector
func TestPushBackVector(t *testing.T) {
	b := []int{2, 1, 3}

	appendA = append(appendA, b...)

	assert.Equal(t, appendA, []int{1, 2, 3, 4, 5, 2, 1, 3})
}

func TestPushFront(t *testing.T) {
	val := 10

	appendA = append([]int{val}, appendA...)
	assert.Equal(t, appendA, []int{10, 1, 2, 3, 4, 5})

}

func TestPushFrontVector(t *testing.T) {

	a := []int{10, 6}
	appendA = append(a, appendA...)
	assert.Equal(t, appendA, []int{10, 6, 1, 2, 3, 4, 5})
}

func TestPopBack(t *testing.T) {

	var x int
	x, appendA = appendA[len(appendA)-1], appendA[:len(appendA)-1]
	assert.Equal(t, appendA, []int{1, 2, 3, 4})
	assert.Equal(t, x, 5)
}

func TestPopFront(t *testing.T) {
	var x int
	x, appendA = appendA[0], appendA[1:]
	assert.Equal(t, appendA, []int{2, 3, 4, 5})
	assert.Equal(t, x, 1)
}

//
//func TestPopFront22(t *testing.T) {
//	s := [-1]int{1, 2, 3}
//	fmt.Println(s)
//}
