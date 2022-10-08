package solution

/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	var row [9][10]int
	var col [9][10]int
	var box [9][10]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			curNum := board[i][j] - '0'
			if row[i][curNum] == 1 || col[j][curNum] == 1 || box[i/3*3+j/3][curNum] == 1 {
				return false
			}
			row[i][curNum] = 1
			col[j][curNum] = 1
			box[i/3*3+j/3][curNum] = 1
		}
	}
	return true
}

// @lc code=end
