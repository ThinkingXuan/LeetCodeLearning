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

	var bfs func() bool
	bfs = func() bool {
		var queue []int
		queue = append(queue, 0)
		for len(queue) > 0 {
			start := queue[0]
			queue = queue[1:]

			if memo[start] { // 是访问过的，跳过
				continue
			}

			memo[start] = true // 未访问过的，记录一下

			for i := start + 1; i <= len(s); i++ {
				prefix := s[start:i]
				if mp[prefix] {
					if i < len(s) {
						queue = append(queue, i)
					} else {
						return true
					}
				}
			}
		}
		return false
	}
	return bfs()
}
