package solution

import "math"

/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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
func maxPathSum(root *TreeNode) int {
	ans := math.MinInt

	var dfs func(*TreeNode) int
	dfs = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}
		leftSum := max(dfs(tn.Left), 0)
		rightSum := max(dfs(tn.Right), 0)
		ans = max(ans, leftSum+rightSum+tn.Val)
		return tn.Val + max(leftSum, rightSum)
	}
	dfs(root)

	return ans
}
func max(a int, b int) int {
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
