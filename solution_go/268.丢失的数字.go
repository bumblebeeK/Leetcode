package solution

import "sort"

/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 丢失的数字
 */

// @lc code=start
func missingNumber(nums []int) int {
	sort.Ints(nums)
	for i, v := range nums {
		if i != v {
			return i
		}
	}
	return len(nums)
}

// @lc code=end
