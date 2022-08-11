package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var cutA = []int{1, 2, 3, 4, 5}

func TestCut(t *testing.T) {
	// 包括1， 不包括3
	i := 1
	j := 3

	cutA = append(cutA[:i], cutA[j:]...)

	assert.Equal(t, cutA, []int{1, 4, 5})
}

func TestCut_GC(t *testing.T) {
	// 包括1， 不包括3
	i := 1
	j := 3

	copy(cutA[i:], cutA[j:])
	for k, n := len(cutA)-j+i, len(cutA); k < n; k++ {
		cutA[k] = 0 // nil or the zero value of T
	}
	cutA = cutA[:len(cutA)-j+i]

	assert.Equal(t, cutA, []int{1, 4, 5})
}
