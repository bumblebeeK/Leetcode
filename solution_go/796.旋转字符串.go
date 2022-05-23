package solution

import "strings"

/*
 * @lc app=leetcode.cn id=796 lang=golang
 *
 * [796] 旋转字符串
 */

// @lc code=start
func rotateString(s string, goal string) bool {
	return len(s) == len(goal) && strings.Contains(strings.Repeat(s, 2), goal)
}

// @lc code=end
