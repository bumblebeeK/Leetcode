package solution

import "math"

/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		for _, v := range coins {
			if v <= i && dp[i] > (dp[i-v]+1) {
				dp[i] = dp[i-v] + 1
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// @lc code=end

func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
		for _, v := range coins {
			if i-v >= 0 && dp[i-v] != math.MaxInt && dp[i] > dp[i-v]+1 {
				dp[i] = dp[i-v] + 1
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
