package solutiongo

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode.cn id=894 lang=golang
 *
 * [894] 所有可能的真二叉树
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
func allPossibleFBT(n int) []*TreeNode {
	if n&1 == 0 {
		return []*TreeNode{}
	}
	if n == 1 {
		return []*TreeNode{
			&TreeNode{
				Val: 0,
			},
		}
	}
	ans := []*TreeNode{}
	for i := 1; i < n; i += 2 {
		left := allPossibleFBT(i)
		right := allPossibleFBT(n - 1 - i)
		for _, l := range left {
			for _, r := range right {
				ans = append(ans, &TreeNode{
					Val:   0,
					Left:  l,
					Right: r,
				})
			}
		}
	}

	return ans
}

// @lc code=end
