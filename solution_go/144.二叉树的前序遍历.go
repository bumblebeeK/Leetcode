package solution

/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	arr := []int{}

	VLR(root, &arr)
	return arr
}

func VLR(node *TreeNode, arr *[]int) {
	if node != nil {
		*arr = append(*arr, node.Val)
		VLR(node.Left, arr)
		VLR(node.Right, arr)
	}
}

// @lc code=end
