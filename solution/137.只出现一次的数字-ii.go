package solution

import "sort"

/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 */

// @lc code=start
func singleNumber(nums []int) int {
	sort.Ints(nums)
	if len(nums) == 1 {
		return nums[0]
	}
	for i := 1; i < len(nums)-1; {
		if nums[i] != nums[i-1] || nums[i] != nums[i+1] {
			if nums[i] == nums[i+1] {
				return nums[i-1]
			}
			if nums[i] == nums[i-1] {
				return nums[i+1]
			}
			return nums[i]
		}
		i += 3
	}
	return nums[len(nums)-1]
}

// @lc code=end
