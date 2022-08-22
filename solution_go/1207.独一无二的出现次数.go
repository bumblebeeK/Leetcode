package solution

/*
 * @lc app=leetcode.cn id=1207 lang=golang
 *
 * [1207] 独一无二的出现次数
 */

// @lc code=start
func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	count := make(map[int]int)
	for _, v := range m {
		count[v]++
	}
	if len(m) != len(count) {
		return false
	}
	return true
}

// @lc code=end
