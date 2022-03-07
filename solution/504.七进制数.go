package solution

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=504 lang=golang
 *
 * [504] 七进制数
 */

// @lc code=start
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	res := ""
	absN, isNegative := abs(num)

	for absN > 0 {
		n := absN % 7   //2
		absN = absN / 7 //14
		res = strconv.Itoa(n) + res
	}
	if isNegative {
		return "-" + res
	}
	return res
}

func abs(num int) (int, bool) {
	if num < 0 {
		return -num, true
	}
	return num, false
}

// @lc code=end
