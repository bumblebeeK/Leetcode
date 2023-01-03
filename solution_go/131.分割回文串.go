package solution

/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start
func partition(s string) [][]string {
	n := len(s)
	ans := [][]string{}
	var dp [][]int8
	dp = make([][]int8, n)
	for i := range dp {
		dp[i] = make([]int8, n)
	}
	var isPalindrome func(i, j int) int8
	isPalindrome = func(i, j int) int8 {
		if i == j || (j-i == 1 && s[i] == s[j]) {
			return 1
		}
		if dp[i][j] != 0 {
			return dp[i][j]
		}
		dp[i][j] = -1
		if s[i] == s[j] {
			dp[i][j] = isPalindrome(i+1, j-1)
		}
		return dp[i][j]
	}
	split := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), split...))
			return
		}
		for j := i; j < n; j++ {
			if isPalindrome(i, j) == 1 {
				split = append(split, s[i:j+1])
				dfs(j + 1)
				split = split[:len(split)-1]
			}
		}
	}
	dfs(0)
	return ans
}

// @lc code=end
