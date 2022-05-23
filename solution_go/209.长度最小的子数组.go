package solution

/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	mark, total, slow := len(nums)+1, 0, 0
	for fast, v := range nums {
		total = total + v
		for total >= target {
			mark = min(fast-slow+1, mark)
			total -= nums[slow]
			slow++
		}
	}
	if mark == len(nums)+1 {
		return 0
	} else {
		return mark
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
