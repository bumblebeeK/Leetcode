package solution

/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = cost[i] + minCost(dp[i-1], dp[i-2])
	}
	return minCost(dp[len(cost)-1], dp[len(cost)-2])
}

func minCost(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
