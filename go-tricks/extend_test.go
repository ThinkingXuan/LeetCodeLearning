package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var extendA = []int{1, 2, 3, 4, 5}

func TestExtend(t *testing.T) {
	j := 2 // 扩展长度
	extendA = append(extendA, make([]int, j)...)

	assert.Equal(t, extendA, []int{1, 2, 3, 4, 5, 0, 0})
}

func TestExtend2(t *testing.T) {
	i := 2 // 从下标i开始，拓展5个长度
	j := 5
	extendA = append(extendA[:i], append(make([]int, j), extendA[i:]...)...)

	assert.Equal(t, extendA, []int{1, 2, 0, 0, 0, 0, 0, 3, 4, 5})

}
