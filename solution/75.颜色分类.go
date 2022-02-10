package solution

/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int) {
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}

	for i, j := 0, 0; i < 3; i++ {
		for m[i] != 0 {
			m[i]--
			nums[j] = i
			j++
		}
	}
}

// @lc code=end
