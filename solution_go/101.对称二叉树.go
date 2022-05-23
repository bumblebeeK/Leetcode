package solution

/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
func isSymmetric(root *TreeNode) bool {
	return checkIsSymmetric(root, root)
}
func checkIsSymmetric(p *TreeNode, q *TreeNode) bool {
	if nil == p && nil == q {
		return true
	}
	if nil == p || nil == q {
		return false
	}
	return p.Val == q.Val && checkIsSymmetric(p.Left, q.Right) && checkIsSymmetric(p.Right, q.Left)
}

// @lc code=end
