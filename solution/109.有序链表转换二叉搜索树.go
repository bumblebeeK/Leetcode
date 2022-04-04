package solution

/*
 * @lc app=leetcode.cn id=109 lang=golang
 *
 * [109] 有序链表转换二叉搜索树
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	var buildTree func(int, int) *TreeNode
	buildTree = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		// 1 2 
		mid := (left + right) >> 1
		node := &TreeNode{Val: arr[mid]}
		node.Left = buildTree(left, mid-1)
		node.Right = buildTree(mid+1, right)
		return node
	}

	return buildTree(0, len(arr)-1)
}

// @lc code=end
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}
