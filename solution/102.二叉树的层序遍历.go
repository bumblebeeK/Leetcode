package solution

/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := make([]*TreeNode, 1)
	q[0] = root
	curLevel := 0
	for len(q) > 0 {
		curLen := len(q)
		for i := 0; i < curLen; i++ {
			node := q[0]
			q = q[1:]
			if len(res) <= curLevel {
				res = append(res, []int{node.Val})
			} else {
				res[curLevel] = append(res[curLevel], node.Val)
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		curLevel++

	}
	return res
}

// @lc code=end
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
