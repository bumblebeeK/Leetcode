/*
 * @lc app=leetcode.cn id=2583 lang=golang
 *
 * [2583] 二叉树中的第 K 大层和
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
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	stack := []*TreeNode{root}
	var heap []int64
	for len(stack) > 0 {
		var val int64
		l := len(stack)
		for l > 0 {
			cur := stack[0]
			val += int64(cur.Val)
			stack = stack[1:]
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			l--
		}
		heap = append(heap, val)
	}
	sort.Slice(heap, func(i, j int) bool {
		return heap[i] > heap[j]
	})
	if k > len(heap) {
		return -1
	}
	return heap[k-1]

}

// @lc code=end

