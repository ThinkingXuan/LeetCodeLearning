package main

//给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。
//
// 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
//
// 能不能dfs看数据量大小，数据量不大时可以剪枝和记忆化搜索。
func wordBreak(s string, wordDict []string) bool {
	mp := map[string]bool{}
	for _, v := range wordDict {
		mp[v] = true
	}

	// 用一个数组，存储计算的结果，数组索引为指针位置，值为计算的结果。下次遇到相同的子问题，直接返回命中的缓存值，就不用调重复的递归。
	// 记忆化结构
	memo := map[int]bool{}

	var dfs func(index int) bool

	dfs = func(index int) bool {
		if index >= len(s) {
			return true
		}

		if res, ok := memo[index]; ok {
			return res
		}

		for i := index; i < len(s); i++ {
			t := s[index : i+1]
			if mp[t] && dfs(i+1) {
				memo[index] = true // 当前递归的结果存一下
				return true
			}
		}
		memo[index] = false // 当前递归的结果存一下
		return false
	}

	return dfs(0)
}
