package solution

/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	p, q := 1, 1
	for i := 0; i < len(nums); i++ {
		ans[i] = p
		p *= nums[i]
	}
	for j := len(nums) - 1; j > 0; j-- {
		q *= nums[j]
		ans[j-1] *= q
	}
	return ans
}

// @lc code=end
