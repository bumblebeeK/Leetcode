package solution

/*
 * @lc app=leetcode.cn id=129 lang=golang
 *
 * [129] 求根节点到叶节点数字之和
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
func sumNumbers(root *TreeNode) int {
	res := 0
	queue := []*TreeNode{root}
	queueVal := []int{root.Val}
	for len(queue) > 0 {
		node := queue[0]
		val := queueVal[0]
		queue = queue[1:]
		queueVal = queueVal[1:]
		if node.Left == nil && node.Right == nil {
			res += val
			continue
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
			queueVal = append(queueVal, val*10+node.Left.Val)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			queueVal = append(queueVal, val*10+node.Right.Val)
		}

	}
	return res
}

// @lc code=end
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
