package solution

import "fmt"

/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	var c2, c3, c5 = 1, 1, 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		n2, n3, n5 := dp[c2]*2, dp[c3]*3, dp[c5]*5
		dp[i] = min(n2, n3, n5)
		if dp[i] == n2 {
			c2++
		}
		if dp[i] == n3 {
			c3++
		}
		if dp[i] == n5 {
			c5++
		}
	}
	return dp[n]
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= c && b <= a {
		return b
	}
	return c
}

// @lc code=end
