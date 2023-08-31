package solution

/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	x := 0
	for i := 0; i < len(t); i++ {
		if x == len(s) {
			return true
		}
		if s[x] == t[i] {
			x++
		}
	}
	return x == len(s)

}

// @lc code=end
