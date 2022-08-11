package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var deleteA = []int{1, 2, 3, 4, 5}

func TestDelete(t *testing.T) {
	i := 2
	deleteA = append(deleteA[:i], deleteA[i+1:]...)

	assert.Equal(t, deleteA, []int{1, 2, 4, 5})
}

func TestDelete2(t *testing.T) {
	i := 2
	deleteA = deleteA[:i+copy(deleteA[i:], deleteA[i+1:])]

	assert.Equal(t, deleteA, []int{1, 2, 4, 5})
}

// 删除某个下标的元素，带GC
func TestDeleteGC(t *testing.T) {
	i := 2
	if i < len(deleteA)-1 {
		copy(deleteA[i:], deleteA[i+1:])
	}
	deleteA[len(deleteA)-1] = 0 // nil or the zero value of T
	deleteA = deleteA[:len(deleteA)-1]

	assert.Equal(t, deleteA, []int{1, 2, 4, 5})
}
