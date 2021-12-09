package solution

import (
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=179 lang=golang
 *
 * [179] 最大数
 */

// @lc code=start
func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		ni, ji := 10, 10
		for ni <= x {
			ni *= 10
		}
		for ji <= y {
			ji *= 10
		}
		return ji*x+y > ni*y+x
	})
	if nums[0] == 0 {
		return "0"
	}
	b := []byte{}
	for _, v := range nums {
		b = append(b, strconv.Itoa(v)...)
	}
	return string(b)
}

// @lc code=end
