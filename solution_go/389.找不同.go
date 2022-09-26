package solution

/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	var diff byte
	for v := range s {
		diff ^= s[v] ^ t[v]
	}
	return diff ^ t[len(t)-1]
}

// @lc code=end
