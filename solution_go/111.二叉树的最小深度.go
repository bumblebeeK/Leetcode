package solution

/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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
func minDepth(root *TreeNode) int {
	depth := 0
	if root != nil {
		depth += 1
		lD := minDepth(root.Left)
		rD := minDepth(root.Right)
		if root.Left == nil {
			return depth + rD
		}
		if root.Right == nil {
			return depth + lD
		}

		depth += min(lD, rD)
	}
	return depth
}

func min(a, b int) int {
	if a < b {
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
