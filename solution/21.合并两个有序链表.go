package solution

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res, head := &ListNode{}, &ListNode{}
	res = head
	for list1 != nil || list2 != nil {
		if list1 == nil {
			head.Next = list2
			return res.Next
		}
		if list2 == nil {
			head.Next = list1
			return res.Next
		}
		if list1.Val > list2.Val {
			head.Next = list2
			head = head.Next
			list2 = list2.Next
		} else {
			head.Next = list1
			head = head.Next
			list1 = list1.Next
		}
	}
	return res.Next
}

// @lc code=end
type ListNode struct {
	Val  int
	Next *ListNode
}
