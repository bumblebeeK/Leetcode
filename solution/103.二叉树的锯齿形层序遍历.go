package solution

/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	q := []*TreeNode{}
	q = append(q, root)
	for len(q) > 0 {
		level := []int{}
		for _, v := range q {
			level = append(level, v.Val)
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
			q = q[1:]
		}
		ans = append(ans, level)
	}
	for i := 1; i < len(ans); i += 2 {
		for j, k := 0, len(ans[i])-1; j < k; j, k = j+1, k-1 {
			ans[i][j], ans[i][k] = ans[i][k], ans[i][j]
		}
	}
	return ans
}

// @lc code=end

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
