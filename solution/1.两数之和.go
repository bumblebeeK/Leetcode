package solution

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		t, ok := m[v]
		if ok {
			return []int{t, i}
		} else {
			m[target-v] = i
		}
	}
	return nil
}

// @lc code=end
