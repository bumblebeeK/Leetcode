package solution

/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start

func oneCount(c int) int {
	one := 0
	for ; c > 0; c &= c - 1 {
		one++
	}
	return one
}

func countBits(n int) []int {
	ret := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ret[i] = oneCount(i)
	}
	return ret
}

// @lc code=end
