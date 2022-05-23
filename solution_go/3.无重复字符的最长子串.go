package solution

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if s == "" || len(s) == 0 {
		return 0
	}
	var m = make(map[byte]int)
	cur, prev := 0, 0
	for i := 0; i < len(s); i++ {
		if index, ok := m[s[i]]; ok && index >= i-cur {
			prev = max(prev, cur)
			cur = i - index

		} else {
			cur++
		}
		m[s[i]] = i
	}
	return max(cur, prev)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
