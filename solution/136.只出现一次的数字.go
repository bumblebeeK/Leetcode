package solution

/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

// @lc code=start
func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res

	// if len(nums) == 1 {
	// 	return nums[0]
	// }
	// sort.Ints(nums)
	// if nums[len(nums)-1] != nums[len(nums)-2] {
	// 	return nums[len(nums)-1]
	// }
	// for i := len(nums) - 2; i > 0; i-- {
	// 	if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
	// 		return nums[i]
	// 	}
	// }
	// return nums[0]
}

// @lc code=end
