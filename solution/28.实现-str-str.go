package solution

/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if haystack == needle || needle == "" {
		return 0
	}
	if len(needle) == 0 {
		return -1
	}

	next := getNext(needle)
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}
	return -1
}

func getNext(needle string) []int {
	next := make([]int, len(needle))
	j := 0
	for i := 1; i < len(needle); i++ {
		for j > 0 && needle[i] != needle[j] {
			j = next[j-1]
		}

		if needle[i] == needle[j] {
			j++
		}
		next[i] = j
	}
	return next
}

// @lc code=end
