/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
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
func getMinimumDifference(root *TreeNode) int {

	pre := math.MaxInt
	ans := math.MaxInt
	var find func(*TreeNode)
	find = func(node *TreeNode) {
		if node == nil {
			return
		}
		find(node.Left)
		ans = min(ans, abs(pre, node.Val))
		pre = node.Val
		find(node.Right)
	}
	find(root)
	return ans
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

