/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	l := len(postorder)
	if l == 0 {
		return nil
	}
	iIdx := -1
	for i := 0; i < l; i++ {
		if inorder[i] == postorder[l-1] {
			iIdx = i
			break
		}
	}
	n := &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  buildTree(inorder[:iIdx], postorder[:iIdx]),
		Right: buildTree(inorder[iIdx+1:], postorder[iIdx:l-1]),
	}
	return n
}

// @lc code=end

