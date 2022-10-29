package solution

/*
 * @lc app=leetcode.cn id=1822 lang=golang
 *
 * [1822] 数组元素积的符号
 */

// @lc code=start
func arraySign(nums []int) int {
	even := 0
	for _, v := range nums {
		if v == 0 {
			return 0
		}
		if v < 0 {
			even++
		}
	}
	if even%2 == 0 {
		return 1
	}
	return -1

}

// @lc code=end
