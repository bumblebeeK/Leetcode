package solution

/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	return Max(DFS(root))
}

func DFS(node *TreeNode) (int, int) {
	if nil == node {
		return 0, 0
	}
	lr, ls := DFS(node.Left)
	rr, rs := DFS(node.Right)
	return node.Val + ls + rs, Max(ls, lr) + Max(rr, rs)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
