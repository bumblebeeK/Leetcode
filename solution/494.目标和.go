package solution

/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {

	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum < target || (sum+target)%2 != 0 {
		return 0
	}
	p := (sum + target) / 2

	dp := make([]int, p+2)
	dp[0] = 1
	for _, v := range nums {
		for i := p; i >= v; i-- {
			dp[i] = dp[i-v] + dp[i]
		}
	}
	return dp[p]
}

// @lc code=end
