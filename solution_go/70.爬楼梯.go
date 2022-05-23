package solution

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	pre, cur := 0, 1
	for i := 0; i < n; i++ {
		tmp := cur
		cur += pre
		pre = tmp
	}
	return cur
}

// @lc code=end

//
