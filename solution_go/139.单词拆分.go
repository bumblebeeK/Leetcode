package solution

/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	dic := make(map[string]bool)
	for _, v := range wordDict {
		dic[v] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(dp); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dic[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// @lc code=end
