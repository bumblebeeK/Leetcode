package solution

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=1154 lang=golang
 *
 * [1154] 一年中的第几天
 */

// @lc code=start
func dayOfYear(date string) int {
	var arrYear = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	dateSplit := strings.Split(date, "-")
	res, _ := strconv.Atoi(dateSplit[2])
	year, _ := strconv.Atoi(dateSplit[0])
	month, _ := strconv.Atoi(dateSplit[1])
	for i := 0; i < month-1; i++ {
		res += arrYear[i]
	}
	if month > 2 {
		return res + isLeap(year)
	}
	return res
}

func isLeap(year int) int {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return 1
	}
	return 0
}

// @lc code=end
