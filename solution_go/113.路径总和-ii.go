package solution

/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	if root.Left == nil && root.Right == nil {
		if targetSum == root.Val {
			return [][]int{{root.Val}}
		} else {
			return [][]int{}
		}
	}
	var res = [][]int{}
	l := pathSum(root.Left, targetSum-root.Val)
	path := []int{root.Val}

	for _, v := range l {
		path = append(path, v...)
		res = append(res, path)
		path = []int{root.Val}
	}
	r := pathSum(root.Right, targetSum-root.Val)
	path = []int{root.Val}

	for _, v := range r {
		path = append(path, v...)
		res = append(res, path)
		path = []int{root.Val}

	}
	return res
}

// @lc code=end
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
