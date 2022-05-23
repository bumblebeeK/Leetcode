package solution

/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Right)
	flatten(root.Left)
	curRight := root.Right
	root.Right = root.Left
	root.Left = nil
	for root.Right != nil {
		root = root.Right
	}
	root.Right = curRight
}

// @lc code=end

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
