package solution

/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	var m map[byte]int = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	if len(s) == 1 {
		return m[s[0]]
	}
	res := 0
	i := 0
	for i = 0; i < len(s)-1; {
		if s[i] == 'I' && (s[i+1] == 'V' || s[i+1] == 'X') {
			res += m[s[i+1]] - m[s[i]]
			i += 2
		} else if s[i] == 'X' && (s[i+1] == 'L' || s[i+1] == 'C') {
			res += m[s[i+1]] - m[s[i]]
			i += 2
		} else if s[i] == 'C' && (s[i+1] == 'D' || s[i+1] == 'M') {
			res += m[s[i+1]] - m[s[i]]
			i += 2
		} else {
			res += m[s[i]]
			i++
		}
	}
	if i == len(s)-1 {
		res += m[s[len(s)-1]]
	}
	return res
}

// @lc code=end
