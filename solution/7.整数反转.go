package solution

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	z := 0
	for x != 0 {
		z = x%10 + z*10
		x = x / 10
	}
	if z > 1<<31-1 || z < -(1<<31) {
		return 0
	}
	return z
}

// @lc code=end
