package solution

/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层序遍历 II
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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var ans [][]int
	q := []*TreeNode{root}
	for len(q) > 0 {
		level := []int{}
		for _, v := range q {
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
			q = q[1:]
			level = append(level, v.Val)
		}
		ans = append(ans, level)
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}

	return ans
}

// @lc code=end

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
