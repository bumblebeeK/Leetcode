package solution

/*
 * @lc app=leetcode.cn id=1790 lang=golang
 *
 * [1790] 仅执行一次字符串交换能否使两个字符串相等
 */

// @lc code=start
func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) == 2 {
		return s1[0] == s2[1] && s1[1] == s2[0]
	}
	var diff0, diff1 byte
	for i, j := 0, 0; i < len(s1); i, j = i+1, j+1 {
		if s1[i] == s2[j] {
			continue
		} else {
			if diff0 == 0 {
				diff0 = s1[i] - s2[j]
				continue
			}
			if diff1 == 0 {
				diff1 = s1[i] - s2[j]
				continue
			}

			return false

		}
	}
	if diff0+diff1 == 0 {
		return true
	}
	return false
}

// @lc code=end
