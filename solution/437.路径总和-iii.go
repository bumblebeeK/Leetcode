package solution

/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := pathSumSol(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res

}

func pathSumSol(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := 0
	if targetSum == root.Val {
		res++
	}
	res += pathSumSol(root.Left, targetSum-root.Val)
	res += pathSumSol(root.Right, targetSum-root.Val)

	return res

}

// @lc code=end
