package solution

/*
 * @lc app=leetcode.cn id=1277 lang=golang
 *
 * [1277] 统计全为 1 的正方形子矩阵
 */

// @lc code=start
func countSquares(matrix [][]int) int {
	count := 0
	row, col := len(matrix), len(matrix[0])
	if row < 0 || col < 0 {
		return 0
	}
	dp := make([][]int, row)

	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
		if matrix[i][0] == 1 {
			count++
			dp[i][0] = 1
		}
	}
	for j := 1; j < col; j++ {
		if matrix[0][j] == 1 {
			count++
			dp[0][j] = 1
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == 1 {
				dp[i][j] = min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
				count += dp[i][j]
			}
		}
	}
	return count
}

func min(a, b, c int) int {
	if a > b {
		b, a = a, b
	}
	if a > c {
		return c
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
