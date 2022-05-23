package solution

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := strs[0]

	for i := 1; i < len(strs); i++ {
		m := min(len(res), len(strs[i]))
		j := 0
		for ; j < m; j++ {
			if res[j] != strs[i][j] {
				break
			}
		}
		res = res[:j]
		if len(res) == 0 {
			return res
		}
	}

	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
