package solution

/*
 * @lc app=leetcode.cn id=1662 lang=golang
 *
 * [1662] 检查两个字符串数组是否相等
 */

// @lc code=start
func arrayStringsAreEqual(word1 []string, word2 []string) bool {

	var b1, b2 []byte
	for _, v := range word1 {
		b1 = append(b1, []byte(v)...)
	}
	for _, v := range word2 {
		b2 = append(b2, []byte(v)...)
	}

	return string(b1) == string(b2)
}

// @lc code=end
