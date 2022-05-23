package solution

/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	for i, v := range nums {
		if (m[v] != 0 || nums[0] == v && i != 0) && k >= Abs(m[v], i) {
			return true
		} else {
			m[v] = i
		}
	}
	return false
}

func Abs(a int, b int) (s int) {
	s = a - b
	if s > 0 {
		return
	} else {
		s = -s
		return
	}
}

// @lc code=end
