package solution

import "fmt"

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	ans := [][]int{}
	if len(nums) == 0 {
		return ans
	}
	var dfs func(res []int, idx int)
	dfs = func(res []int, idx int) {
		if idx == len(nums) {
			fmt.Println(nums)
			s := make([]int, len(nums))
			copy(s, nums)
			ans = append(ans, s)
		}
		for i := idx; i < len(nums); i++ {
			nums[idx], nums[i] = nums[i], nums[idx]
			dfs(res, idx+1)
			nums[idx], nums[i] = nums[i], nums[idx]
		}
	}
	dfs(nums, 0)
	return ans
}

// @lc code=end
