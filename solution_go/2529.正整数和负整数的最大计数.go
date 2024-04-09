package solutiongo

/*
 * @lc app=leetcode.cn id=2529 lang=golang
 *
 * [2529] 正整数和负整数的最大计数
 */

// @lc code=start
func maximumCount(nums []int) int {
	l := len(nums)
	neg, pos := 0, 0
	for i, j := 0, l-1; i <= j; {
		if nums[i] >= 0 && nums[j] <= 0 {
			break
		}
		if nums[i] < 0 {
			neg++
			i++
		}
		if nums[j] > 0 {
			pos++
			j--
		}
	}
	return max(neg, pos)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end
