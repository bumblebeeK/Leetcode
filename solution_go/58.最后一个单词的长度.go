package solution

/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	l := 0
	lenS := len(s)
	if lenS == 0 {
		return 0
	}
	for i := lenS - 1; i >= 0; i-- {
		if s[i] == 32 && l != 0 {
			return l
		} else if s[i] != 32 {
			l++
		}
	}
	return l
}

// @lc code=end
