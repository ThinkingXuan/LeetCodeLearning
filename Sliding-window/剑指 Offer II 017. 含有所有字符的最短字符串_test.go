package main

import "testing"

//https://leetcode-cn.com/problems/M1oyTv/
/*
给定两个字符串 s 和t 。返回 s 中包含t的所有字符的最短子字符串。如果 s 中不存在符合条件的子字符串，则返回空字符串 "" 。

如果 s 中存在多个符合条件的子字符串，返回任意一个。

注意： 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
*/

// 滑动窗口
func minWindow(s string, t string) string {
	sMap, tMap := map[byte]int{}, map[byte]int{}

	sLen, tLen := len(s), len(t)

	for i := 0; i < tLen; i++ {
		tMap[t[i]]++
	}

	res := sLen + 1
	left, right := 0, -1

	for i, j := 0, 0; j < sLen; j++ {
		sMap[s[j]]++

		for isEqual(sMap, tMap) {
			if res > j-i+1 {
				res = j - i + 1
				left = i
				right = j
			}
			sMap[s[i]]--
			i++
		}
	}
	return s[left : right+1]
}

func isEqual(sMap map[byte]int, tMap map[byte]int) bool {
	for key, value := range tMap {
		if sMap[key] < value {
			return false
		}
	}
	return true
}

func TestMinWindow(t *testing.T) {
	testCase := []struct{ in1, int2, want string }{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
	}

	for _, tc := range testCase {
		res := minWindow(tc.in1, tc.int2)
		if res != tc.want {
			t.Errorf("MinWindow： %v, want to %v", res, tc.want)
		}
	}
}
