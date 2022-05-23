package solution

/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	node := &TreeNode{Val: preorder[0]}
	t := 0
	for i, v := range inorder {
		if v == preorder[0] {
			t = i
		}
	}
	node.Left = buildTree(preorder[1:t+1], inorder[:t])
	node.Right = buildTree(preorder[t+1:], inorder[t+1:])
	return node

}

// @lc code=end
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
