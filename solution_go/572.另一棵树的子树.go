package solution

/*
 * @lc app=leetcode.cn id=572 lang=golang
 *
 * [572] 另一棵树的子树
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
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil && root == nil {
		return true
	}
	if (subRoot == nil && root != nil) || (subRoot != nil && root == nil) {
		return false
	}
	return isSameTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(leftT *TreeNode, rightT *TreeNode) bool {
	if leftT == nil && rightT == nil {
		return true
	}

	if (leftT == nil && rightT != nil) || (leftT != nil && rightT == nil) {
		return false
	}
	return isSameTree(leftT.Left, rightT.Left) && leftT.Val == rightT.Val && isSameTree(leftT.Right, rightT.Right)

}

// @lc code=end

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
