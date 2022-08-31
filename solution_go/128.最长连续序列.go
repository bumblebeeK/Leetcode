package solution

/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	longest := 0
	m := map[int]bool{}
	for _, v := range nums {
		m[v] = true
	}

	for v := range m {
		if !m[v-1] {
			curNum := v
			curMax := 1
			for m[curNum+1] {
				curMax++
				curNum++
			}
			if curMax > longest {
				longest = curMax
			}
		}
	}
	return longest
}

// @lc code=end
