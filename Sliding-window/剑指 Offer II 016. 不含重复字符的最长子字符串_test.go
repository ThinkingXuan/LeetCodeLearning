package main

import "testing"

//剑指 Offer II 016. 不含重复字符的最长子字符串
//https://leetcode-cn.com/problems/wtcaE1/
//给定一个字符串 s ，请你找出其中不含有重复字符的 最长连续子字符串 的长度。

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	lookup := make(map[byte]int)
	maxStr := 0
	left := 0

	for i := 0; i < len(s); i++ {

		if _, ok := lookup[s[i]]; ok {
			left = max(left, lookup[s[i]]+1)
		}

		maxStr = max(maxStr, i-left+1)
		lookup[s[i]] = i
	}
	return maxStr
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func TestLengthOfLongestSubstring(t *testing.T) {
	testCase := []struct {
		in   string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
	}

	for _, tc := range testCase {
		res := lengthOfLongestSubstring(tc.in)
		if res != tc.want {
			t.Errorf("TestLengthOfLongestSubstring: %v, want %v", res, tc.want)
		}
	}
}
