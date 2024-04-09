package solutiongo

import "fmt"

/*
 * @lc app=leetcode.cn id=228 lang=golang
 *
 * [228] 汇总区间
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	var ans []string
	if len(nums) == 0 {
		return ans
	}

	start := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 {
			continue
		}
		if start != i-1 {
			ans = append(ans, fmt.Sprintf("%d->%d", nums[start], nums[i-1]))
		} else {
			ans = append(ans, fmt.Sprintf("%d", nums[start]))
		}
		start = i
	}
	if start != len(nums)-1 {
		ans = append(ans, fmt.Sprintf("%d->%d", nums[start], nums[len(nums)-1]))
	} else {
		ans = append(ans, fmt.Sprintf("%d", nums[start]))
	}

	return ans
}

// @lc code=end
