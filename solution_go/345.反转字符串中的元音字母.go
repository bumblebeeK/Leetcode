package solution

/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 */

// @lc code=start
func reverseVowels(s string) string {
	if len(s) == 1 {
		return s
	}
	bytes := []byte(s)
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	left, right := 0, len(bytes)-1
	for left < right {
		isVowels := false
		for _, v := range vowels {
			if v == bytes[left] {
				isVowels = true
			}
		}
		if isVowels {
		right:
			for left < right {
				for _, v := range vowels {
					if v == bytes[right] {
						bytes[left], bytes[right] = bytes[right], bytes[left]
						right--
						break right
					}
				}
				right--
			}
		}

		left++
	}
	return string(bytes)
}

// @lc code=end
