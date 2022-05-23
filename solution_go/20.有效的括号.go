package solution

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	var stack []string
	rm := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	for _, v := range s {
		if r, ok := rm[string(v)]; ok {
			if len(stack) > 0 && r == stack[len(stack)-1:][0] {
				stack = stack[0 : len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, string(v))
		}
	}
	return len(stack) == 0
}

// @lc code=end
