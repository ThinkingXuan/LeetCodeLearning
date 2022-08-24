package main

//字典 wordList 中从单词 beginWord 和 endWord 的 转换序列 是一个按下述规格形成的序列
// beginWord -> s1 -> s2 -> ... -> sk：
//
//
// 每一对相邻的单词只差一个字母。
//
// 对于 1 <= i <= k 时，每个
// si 都在
// wordList 中。注意， beginWord 不需要在
// wordList 中。
//
// sk == endWord
//
//
// 给你两个单词 beginWord 和 endWord 和一个字典 wordList ，返回 从 beginWord 到 endWord 的 最短转换序列
//中的 单词数目 。如果不存在这样的转换序列，返回 0 。
//
// 示例 1：
//
//
//输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot",
//"log","cog"]
//输出：5
//解释：一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog", 返回它的长度 5。
//
//
// 示例 2：
//
//
//输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot",
//"log"]
//输出：0
//解释：endWord "cog" 不在字典中，所以无法进行转换。
//

// BFS + 状态枚举
func ladderLength(beginWord string, endWord string, wordList []string) int {

	wordMap := map[string]bool{}
	for _, r := range wordList {
		wordMap[r] = true
	}

	// endmap不在wordlist中
	if !wordMap[endWord] {
		return 0
	}

	get := func(word string) []string {

		sb := []byte(word)
		var ans []string

		for i := 0; i < len(sb); i++ {
			b := sb[i]
			for j := 0; j < 26; j++ {
				sb[i] = byte('a' + j)
				if wordMap[string(sb)] {
					ans = append(ans, string(sb))
				}
			}
			sb[i] = b
		}
		return ans
	}

	type pair struct {
		status string
		step   int
	}

	var queue []pair
	seen := map[string]bool{}

	queue = append(queue, pair{beginWord, 1})

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, q := range get(cur.status) {
			if !seen[q] {
				if q == endWord {
					return cur.step + 1
				}
				seen[q] = true
				queue = append(queue, pair{q, cur.step + 1})
			}
		}
	}

	return 0
}
