package solution

/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	return Max(goRob(nums[:len(nums)-1]), goRob(nums[1:]))
}

func goRob(nums []int) int {
	cur, pre := 0, 0
	for _, v := range nums {
		temp := Max(pre+v, cur)
		pre = cur
		cur = temp
	}
	return cur
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
