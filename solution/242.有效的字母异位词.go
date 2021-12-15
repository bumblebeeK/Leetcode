package solution

/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	m := map[rune]int{}
	counter := len(s)
	for _, v := range s {
		m[v] += 1
	}
	for _, v := range t {
		if _, ok := m[v]; ok {
			m[v] -= 1
			counter--
			if m[v] < 0 {
				return false
			}
		} else {
			return false
		}
	}
	if counter > 0 {
		return false
	}
	return true
}

// @lc code=end

// func isAnagram(s, t string) bool {
// 	var c1, c2 [26]int
// 	for _, ch := range s {
// 		c1[ch-'a']++
// 	}
// 	for _, ch := range t {
// 		c2[ch-'a']++
// 	}
// 	return c1 == c2
// }
