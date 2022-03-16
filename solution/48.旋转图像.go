package solution

/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start
func rotate(matrix [][]int) {
	n := len(matrix)
	newMatrix := make([][]int, n)
	for i := range newMatrix {
		newMatrix[i] = make([]int, n)
	}
	for x, v := range matrix {
		for y, t := range v {
			newMatrix[y][n-x-1] = t
		}
	}
	copy(matrix, newMatrix)
}

// @lc code=end
