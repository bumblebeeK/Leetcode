package solution

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	prev, curv := 0, 1
	for i := 0; i < n; i++ {
		tmp := curv
		curv += prev
		prev = tmp
	}
	return curv

}

// @lc code=end

//
