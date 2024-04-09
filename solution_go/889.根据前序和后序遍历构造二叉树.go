/*
 * @lc app=leetcode.cn id=889 lang=golang
 *
 * [889] 根据前序和后序遍历构造二叉树
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
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	l := len(postorder)
	if l == 0 {
		return nil
	}
	cur := &TreeNode{
		Val: postorder[l-1],
	}
	idx := -1
	if l == 1 {
		return cur
	}

	for i, v := range postorder {
		if v == preorder[1] {
			idx = i + 1
			break
		}
	}
	cur.Left = constructFromPrePost(preorder[1:1+idx], postorder[:idx])
	cur.Right = constructFromPrePost(preorder[1+idx:], postorder[idx:l-1])

	return cur

}

// @lc code=end

