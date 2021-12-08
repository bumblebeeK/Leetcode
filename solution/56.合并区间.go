package solution

import "sort"

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		} else {
			return intervals[i][0] < intervals[j][0]
		}
	})
	res := [][]int{}
	for _, v := range intervals {
		if len(res) == 0 || res[len(res)-1:][0][1] < v[0] {
			res = append(res, v)
		} else if res[len(res)-1:][0][1] < v[1] {
			res[len(res)-1:][0] = []int{res[len(res)-1:][0][0], v[1]}
		}
	}
	return res
}

// @lc code=end
