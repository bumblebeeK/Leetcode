package leetcode

import "strconv"

/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
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
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	res := []string{}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	l := binaryTreePaths(root.Left)
	for _, v := range l {
		res = append(res, strconv.Itoa(root.Val)+"->"+v)
	}
	r := binaryTreePaths(root.Right)
	for _, v := range r {
		res = append(res, strconv.Itoa(root.Val)+"->"+v)
	}
	return res
}

// @lc code=end
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
