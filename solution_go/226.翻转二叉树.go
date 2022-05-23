package solution

/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
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
// func invertTree(root *TreeNode) *TreeNode {
// 	if nil == root {
// 		return nil
// 	}
// 	tmp := root.Left
// 	root.Left = root.Right
// 	root.Right = tmp
// 	invertTree(root.Left)
// 	invertTree(root.Right)
// 	return root
// }
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var q []*TreeNode
	q = append(q, root)
	for len(q) != 0 {
		element := q[0]
		q = q[1:]
		if element != nil {
			tmp := element.Left
			element.Left = element.Right
			element.Right = tmp
			q = append(q, element.Left)
			q = append(q, element.Right)
		}
	}
	return root
}

// @lc code=end

