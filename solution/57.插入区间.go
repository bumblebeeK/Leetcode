package solution

/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	left, right := newInterval[0], newInterval[1]
	merged := false
	ans := [][]int{}
	for _, interval := range intervals {
		if interval[0] > right {
			if !merged {
				ans = append(ans, []int{left, right})
				merged = true
			}
			ans = append(ans, interval)
		} else if interval[1] < left {
			ans = append(ans, interval)
		} else {
			if left > interval[0] {
				left = interval[0]
			}
			if right < interval[1] {
				right = interval[1]
			}
		}
	}
	if !merged {
		ans = append(ans, []int{left, right})
	}
	return ans
}

// @lc code=end
