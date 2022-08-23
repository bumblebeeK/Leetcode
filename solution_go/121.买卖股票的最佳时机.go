package solution

import "math"

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	var res int
	minPrice := math.MaxInt
	for _, v := range prices {
		if v < minPrice {
			minPrice = v
		}
		res = max(res, v-minPrice)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
