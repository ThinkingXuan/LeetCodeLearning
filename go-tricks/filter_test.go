package main

import "testing"

var filterA = []int{1, 2, 4, 4, 5, 4, 2}

func TestFilter(t *testing.T) {
	n := 0
	for _, x := range filterA {
		if keep(x) {
			a[n] = x
			n++
		}
	}
	a = a[:n]
}

func keep(x int) bool {

	return false
}
