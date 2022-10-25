package solution

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=973 lang=golang
 *
 * [973] 最接近原点的 K 个点
 */

// @lc code=start
func kClosest(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		return (math.Pow(float64(points[i][0]), 2) + math.Pow(float64(points[i][1]), 2)) < (math.Pow(float64(points[j][0]), 2) + math.Pow(float64(points[j][1]), 2))
	})
	return points[:k]
}

// @lc code=end
