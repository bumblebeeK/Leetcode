package solution

import "fmt"

/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
var direction = [][]int{
	[]int{-1, 0},
	[]int{0, 1},
	[]int{0, -1},
	[]int{1, 0},
}

func exist(board [][]byte, word string) bool {
	var m,n int
	if m =len(board); m ==0  {
		return false
	}
	n = len(board[0])
	mark := make([][]bool, m)
	for i := 0; i < m; i++ {
		mark[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				mark[i][j] = true
				if traceBack(i,j,mark,board,word[1:]) == true {
					return true
				}
				mark[i][j] =false
			}
		}
	}
	return false
}


func traceBack(i,j int,mark [][]bool,board [][]byte,word string) bool{
	if len(word) == 0 {
		return true
	}
	for _, v := range direction {
		curI,curJ := i +v[0],j+v[1]
		if curI >=0 && curI < len(board) && curJ >=0 && curJ <len(board[0]) && word[0] == board[curI][curJ] && !mark[curI][curJ] {
			mark[curI][curJ] = true
			if traceBack(curI,curJ,mark,board,word[1:]) == true {
				return true
			}
			mark[curI][curJ] =false
		}
	}
	return false
}



// @lc code=end
