package solution

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		start, end := i+1, len(nums)-1
		for end > start {
			sum := nums[i] + nums[start] + nums[end]
			if simpleAbs(sum-target) < simpleAbs(ans-target) {
				ans = sum
			}
			if sum > target {
				end--
			} else if sum < target {
				start++
			} else {
				return target
			}
		}
	}
	return ans
}

func simpleAbs(num int) int {
	if num > 0 {
		return num
	} else {
		return -num
	}
}

// @lc code=end
