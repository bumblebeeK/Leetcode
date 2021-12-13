package solution

/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(dfs(root.Left), dfs(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(dfs(node.Left), dfs(node.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a, b int) (v int) {
	if v = a - b; v > 0 {
		return
	}
	return -v
}

// @lc code=end
