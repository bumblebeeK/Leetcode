package solution

/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	// return strings.Index(haystack, needle)
	if len(needle) == 0 {
		return 0
	}
	next := kmpNext(needle)
	for i, j := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != haystack[j] {
			j = next[j]
		}
	}

	return -1

}

func kmpNext(subArray string) []int {
	next := make([]int, len(subArray))
	next[0] = 0
	for i, j := 1, 0; i < len(subArray); i++ {
		for j > 0 && subArray[i] != subArray[j] {
			j = next[j-1]
		}
		if subArray[i] == subArray[j] {
			j++
		}
		next[i] = j
	}
	return next

}

// @lc code=end
