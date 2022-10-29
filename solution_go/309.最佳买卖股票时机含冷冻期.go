package solution

/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
func maxProfit(prices []int) int {
	f1 := 0          // 不持有股票，且不处在冷冻期
	f2 := -prices[0] // 持有股票
	f3 := 0          // 不持有股票，且处在冷冻期
	for i := 1; i < len(prices); i++ {
		newf1 := max(f1, f3)
		newf2 := max(f1-prices[i], f2)
		newf3 := f2 + prices[i]
		f1 = newf1
		f2 = newf2
		f3 = newf3
	}
	return max3(f1, f2, f3)
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func max3(n1, n2, n3 int) int {
	if n1 > n2 {
		n2 = n1
	}
	if n2 > n3 {
		n3 = n2
	}
	return n3
}

// @lc code=end
