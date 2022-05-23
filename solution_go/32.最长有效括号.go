package solution

/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
func longestValidParentheses(s string) int {
	mark := make([]int, len(s))
	stack := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == 40 {
			stack = append(stack, i)
		} else {
			if len(stack) == 0 {
				mark[i] = 1
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	for _, v := range stack {
		mark[v] = 1
	}
	cur, max := 0, 0
	for _, v := range mark {
		if v == 1 {
			cur = 0
		} else {
			cur++
			if cur > max {
				max = cur
			}
		}
	}
	if cur > max {
		return cur
	}
	return max
}

// @lc code=end
