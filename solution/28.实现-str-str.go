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
	if len(haystack) == 0 {
		return -1
	}

	return KMP2(needle, haystack)
}

func KMP2(needle string, haystack string) int {
	next := getNext2(needle)
	j, i := 0, 0
	for i < len(haystack) && j < len(needle) {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}

	}
	if len(needle) == j {
		return i - len(needle)
	}

	return -1
}

func KMP1(needle string, haystack string) int {
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
func KMP3(p string, s string) int {

	i, j := 0, 0
	pLen := len(p)
	sLen := len(s)
	next := getNext3(p)
	for i < sLen && j < pLen {
		if j == -1 || s[i] == p[j] { //s[i]!=s[0]=>j=next[0]=-1,第0位不匹配所以i++，j++;j=0
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == pLen {
		return i - j
	} else {
		return -1
	}

}
func getNext2(needle string) []int {
	next := make([]int, len(needle)+1, len(needle)+1)
	next[0] = -1
	next[1] = 0
	j, i := 0, 1
	for i < len(needle)-1 {
		if j == -1 || needle[i] == needle[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}

	}
	return next
}

func getNext3(p string) []int {
	pLen := len(p)
	next := make([]int, pLen+1, pLen+1)
	next[0] = -1
	next[1] = 0
	i := 0
	j := 1
	for j < pLen-1 { //因为next[pLen-1]由s[i] == s[pLen-2]算出
		if i == -1 || p[i] == p[j] { //-1代表了起始位不匹配，i=0,s[0]!=s[j]=>i=next[0]=-1
			i++
			j++
			if p[i] != p[j] { //因为出现在j位置不匹配的话会跳到next[j]=i位置去匹配,p[i] == p[j]肯定又是不匹配（优化核心点）
				next[j] = i
			} else {
				next[j] = next[i]
			}

		} else {
			i = next[i]
		}
	}

	return next

}

// @lc code=end
