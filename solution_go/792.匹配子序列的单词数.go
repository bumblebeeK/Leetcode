package solution

import "sort"

/*
 * @lc app=leetcode.cn id=792 lang=golang
 *
 * [792] 匹配子序列的单词数
 */

// @lc code=start
func numMatchingSubseq(s string, words []string) int {
	pair := [26][]int{}
	for i, v := range s {
		pair[v-'a'] = append(pair[v-'a'], i)
	}
	ans := len(words)
	for _, w := range words {
		if len(w) > len(s) {
			ans--
			continue
		}
		p := -1
		for _, c := range w {
			ps := pair[c-'a']
			i := sort.SearchInts(ps, p+1)
			if i == len(ps) {
				ans--
				break
			}
			p = ps[i]
		}
	}
	return ans
}

// @lc code=end
