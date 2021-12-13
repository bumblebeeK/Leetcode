package solution

/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(n int) int {
	prev, cur := 0, 1
	for n > 0 {
		tmp := cur + prev
		prev = cur
		cur = tmp
		n--
	}
	return prev
}

// @lc code=end
