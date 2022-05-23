package solution

/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */

// @lc code=start
func isUgly(n int) bool {
	if n < 1 {
		return false
	}
	arrUgly := []int{2, 3, 5}
	for _, v := range arrUgly {
		for n%v == 0 {
			n /= v
		}
	}
	return 1 == n
}

// @lc code=end
