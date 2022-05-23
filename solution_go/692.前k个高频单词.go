package solution

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=692 lang=golang
 *
 * [692] 前K个高频单词
 */

// @lc code=start
func topKFrequent(words []string, k int) []string {
	dic := map[string]int{}
	for _, v := range words {
		dic[v] += 1
	}
	arr := make([]string, 0, len(dic))
	for v := range dic {
		arr = append(arr, v)
	}
	fmt.Println(arr)
	sort.Slice(arr, func(i, j int) bool {
		a, b := arr[i], arr[j]
		fmt.Println(arr[i], arr[j])
		return dic[a] > dic[b] || dic[a] == dic[b] && a < b
	})
	return arr[:k]

}

// @lc code=end
