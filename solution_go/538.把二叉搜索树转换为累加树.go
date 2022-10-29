package solution

/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
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
func convertBST(root *TreeNode) *TreeNode {
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
