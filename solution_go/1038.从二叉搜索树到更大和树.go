package solution

/*
 * @lc app=leetcode.cn id=1038 lang=golang
 *
 * [1038] 从二叉搜索树到更大和树
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
func bstToGst(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode)
	val := 0
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Right)
			val += node.Val
			node.Val = val
			dfs(node.Left)
		}
	}
	dfs(root)
	return root
}

// @lc code=end
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
