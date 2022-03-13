package solution

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	ln := len(height)
	res := min(height[ln-1], height[0]) * (ln - 1)
	for i, j := 0, ln-1; i < j; {
		if j-i == 1 {
			return res
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
		res = max(res, (j-i)*min(height[i], height[j]))
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
