package solution

/*
 * @lc app=leetcode.cn id=897 lang=golang
 *
 * [897] 递增顺序搜索树
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
func increasingBST(root *TreeNode) *TreeNode {
	var dummyHead = &TreeNode{-1, nil, nil}
	nowHead := dummyHead
	var dfs func(*TreeNode)
	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		dfs(tn.Left)
		nowHead.Right = tn
		tn.Left = nil
		nowHead = tn
		dfs(tn.Right)
	}
	dfs(root)
	return dummyHead.Right
}

// @lc code=end

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
