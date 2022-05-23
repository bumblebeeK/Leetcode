package solution

/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
// func rob(nums []int) int {
func rob(nums []int) int {
	cur, pre := 0, 0
	for _, v := range nums {
		tmp := cur
		cur = Max(pre+v, cur)
		pre = tmp
	}
	return cur
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
