package solution

/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	curv, prev := 0, 0
	for _, v := range nums {
		temp := curv
		curv = Max(temp, prev+v)
		prev = temp
	}
	return curv
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
