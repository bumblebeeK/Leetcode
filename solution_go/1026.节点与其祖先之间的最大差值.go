package solutiongo

/*
 * @lc app=leetcode.cn id=1026 lang=golang
 *
 * [1026] 节点与其祖先之间的最大差值
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
func maxAncestorDiff(root *TreeNode) int {
	var findMax func(node *TreeNode, ma, mi int) int
	findMax = func(node *TreeNode, ma, mi int) int {
		if node == nil {
			return 0
		}

		diff := max(abs(node.Val, ma), abs(node.Val, mi))
		mi = min(mi, node.Val)
		ma = max(ma, node.Val)

		diff = max(diff, findMax(node.Left, ma, mi))
		diff = max(diff, findMax(node.Right, ma, mi))

		return diff
	}

	return findMax(root, root.Val, root.Val)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// @lc code=end
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
