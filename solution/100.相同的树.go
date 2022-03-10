package solution

/*
 * @lc app=leetcode.cn id=100 lang=golang
 *
 * [100] 相同的树
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
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
	}
	if (p == nil && q != nil) || (q == nil && p != nil) {
		return false
	}
	if q == nil && p == nil {
		return true
	}
	lR, rR := false, false
	if p.Left != nil && q.Left != nil {
		lR = isSameTree(p.Left, q.Left)
	} else if p.Left == nil && q.Left == nil {
		lR = true
	}
	if p.Right != nil && q.Right != nil {
		rR = isSameTree(p.Right, q.Right)
	} else if p.Right == nil && q.Right == nil {
		rR = true
	}

	return lR && rR
}

// @lc code=end

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
