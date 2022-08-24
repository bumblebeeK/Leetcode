package solution

/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	diff := x ^ y
	count := 0
	for diff > 0 {
		diff &= diff - 1
		count++
	}
	return count
}

// @lc code=end
