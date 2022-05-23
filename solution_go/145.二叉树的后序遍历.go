package solution

/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) []int {
	arr := []int{}
	LRD(root, &arr)
	return arr
}

func LRD(node *TreeNode, arr *[]int) {
	if node != nil {
		LRD(node.Left, arr)
		LRD(node.Right, arr)
		*arr = append(*arr, node.Val)
	}
}

// @lc code=end
