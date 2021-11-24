package solution

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	var res []int
	LDR(&res, root)
	return res
}

func LDR(n *[]int, root *TreeNode) {
	if root != nil {
		LDR(n, root.Left)
		*n = append(*n, root.Val)
		LDR(n, root.Right)
	}
}

// @lc code=end

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
