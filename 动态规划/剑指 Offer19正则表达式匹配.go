package main

//请实现一个函数用来匹配包含'. '和'*'的正则表达式。模式中的字符'.'表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但与"aa.a"和"ab*a"均不匹配。

// 动态规划
// s[:i] 表示前i个字符， p[:j]前j个字符

func main() {

}

func isMatch(s string, p string) bool {
	m, n := len(s)+1, len(p)+1
	dp := make([][]bool, m)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	dp[0][0] = true

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {

			if p[j-1] == '*' {
				if dp[i][j-2] {
					dp[i][j] = true
				}

				/////

				/////
			}else {
				if dp[i-1][j-1] && s[i-1] == p[j-1] {
					dp[i][j] = true
				}
				if dp[i-1][j-1] &&  p[j-1] == '.' {
					dp[i][j] = true
				}
			}
		}
	}
	return dp[m-1][n-1]
}
