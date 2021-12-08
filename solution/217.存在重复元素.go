package solution

/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	m := map[int]bool{}
	for _, v := range nums {
		if m[v] {
			return true
		} else {
			m[v] = true
		}
	}
	return false
}

// @lc code=end
