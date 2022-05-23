package solution

import "golang.org/x/tools/go/callgraph/cha"

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	ln := len(s)
	if ln < 2 {
		return s
	}
	maxL := 1
	begin := 0
	dp := make([][]bool, ln)
	for i, _ := range dp {
		dp[i] = make([]bool, ln)
	}
	for i := 0; i < ln; i++ {
		dp[i][i] = true
	}
	for i := 2; i <= ln; i++ {
		for j := 0; j < ln; j++ {
			x := i + j - 1
			if x >= ln {
				break
			}
			
			if s[j] != s[x] {
				
			}
		}
	}
}

// @lc code=end
