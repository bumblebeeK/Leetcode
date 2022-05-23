package solution

import "strings"

/*
 * @lc app=leetcode.cn id=720 lang=golang
 *
 * [720] 词典中最长的单词
 */

// @lc code=start
func longestWord(words []string) string {
	ans := ""
	set := map[string]struct{}{}
	for _, v := range words {
		set[v] = struct{}{}
	}
	for v := range set {
		curLen, ansLen := len(v), len(ans)
		ilegal := true
		if curLen < ansLen {
			continue
		}
		if curLen == ansLen && strings.Compare(v, ans) != -1 {
			continue
		}
		for i := 1; i <= len(v); i++ {
			if _, ok := set[v[0:i]]; !ok {
				ilegal = false
				break
			}
		}
		if ilegal {
			ans = v
		}
	}
	return ans
}

// @lc code=end
