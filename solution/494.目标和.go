package solution

/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	if target > total || (target+total)%2 == 1 || (target+total) < 0 {
		return 0
	}
	p := (target + total) / 2 //p 为前面为+部分的和
	dp := make([]int, p+1)
	dp[0] = 1
	l := len(nums)
	for _, n := range nums {
		l--
		for i := p; i >= n; i-- {
			dp[i] += dp[i-n]
			if l == 0 {
				return dp[p]
			}
		}
	}
	return dp[p]
}

// @lc code=end
