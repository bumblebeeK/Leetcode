package solution

/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		left, right := 0, 0
		if node.Left != nil {
			left = dfs(node.Left)
		}
		if node.Right != nil {
			right = dfs(node.Right)
		}
		ans = max(left+right, ans)
		return max(left, right) + 1
	}
	dfs(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
