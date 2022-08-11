package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var insertA = []int{1, 2, 3, 4, 5}

func TestInsert(t *testing.T) {

	// 在下标i中插入 val
	i := 2
	val := 10

	insertA = append(insertA[:i], append([]int{val}, insertA[i:]...)...)

	assert.Equal(t, insertA, []int{1, 2, 10, 3, 4, 5})
}
