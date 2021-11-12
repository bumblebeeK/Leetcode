package solution

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	// s = preProcess(s)
	// sL := len(s)
	j := 0

	return ""
}

func preProcess(s string) string {
	if len(s) == 0 {
		return "^$"
	}
	t := "^"
	for i := 0; i < len(s); i++ {
		t += string(s[i])
	}
	t += "#$"
	return t
}

func getNEXT(needele string) []int {
	next := make([]int, len(needele)+1)
	j := 0
	max, index := 0, 0
	next[0] = -1
	next[1] = 0
	for i := 1; i < len(needele); i++ {
		if j > -1 && needele[i] == needele[j] {
			i++
			j++
			next[i] = j
			if j > max {
				max = j
				index = i
			}
		} else {
			j = next[j]
		}
	}
	return []int{max, index}
}

// @lc code=end
