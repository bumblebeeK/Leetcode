package solution

/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	maxSide := 0
	row, col := len(matrix), len(matrix[0])
	dp := make([][]int, row)
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			maxSide = 1
		}
	}
	for i := 0; i < col; i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
			maxSide = 1
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				}

				maxSide = max(maxSide, dp[i][j])
			}
		}
	}
	return maxSide * maxSide
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
