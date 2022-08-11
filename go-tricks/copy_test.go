package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var a = []int{1, 2, 3, 4, 5}

func TestCopy1(t *testing.T) {
	b := make([]int, len(a))
	copy(b, a)

	assert.Equal(t, b, a)
}

func TestCopy2(t *testing.T) {
	var b []int
	b = append([]int(nil), a...)
	assert.Equal(t, b, a)
}

func TestCopy3(t *testing.T) {
	var b []int
	b = append(a[:0:0], a...)
	assert.Equal(t, b, a)
}
