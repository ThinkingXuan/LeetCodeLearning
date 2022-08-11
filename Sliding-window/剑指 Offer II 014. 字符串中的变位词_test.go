package main

import "testing"

//剑指 Offer II 014. 字符串中的变位词
//https://leetcode-cn.com/problems/MPnaiL/
//给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的某个变位词。
//
//换句话说，第一个字符串的排列之一是第二个字符串的 子串 。

func checkInclusion(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)

	if m > n {
		return false
	}

	var cnt1, cnt2 [26]int

	for i, ch := range s1 {
		cnt1[ch-'a']++
		cnt2[s2[i]-'a']++
	}

	if cnt1 == cnt2 {
		return true
	}

	for i := m; i < n; i++ {
		cnt2[s2[i]-'a']++
		cnt2[s2[i-m]-'a']--
		if cnt1 == cnt2 {
			return true
		}
	}

	return false
}

func TestCheckInclusion(t *testing.T) {
	testCase := []struct {
		in1, in2 string
		want     bool
	}{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
	}

	for _, tc := range testCase {
		res := checkInclusion(tc.in1, tc.in2)
		if tc.want != res {
			t.Errorf("checkInclusion: %v, want %v", res, tc.want)
		}
	}
}
