package solution

/*
 * @lc app=leetcode.cn id=915 lang=golang
 *
 * [915] 分割数组
 */

// @lc code=start
func partitionDisjoint(nums []int) int {
	n := len(nums)

	min := make([]int, n)

	min[n-1] = nums[n-1]

	for i := n - 2; i >= 0; i-- {
		min[i] = Min(min[i+1], nums[i])
	}

	for i, max := 0, 0; i < n-1; i++ {
		max = Max(max, nums[i])
		if max <= min[i+1] {
			return i + 1
		}
	}

	return -1
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
