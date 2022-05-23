package solution

/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	ln := 0
	slow := &ListNode{0, head}
	fast := slow
	for fast.Next != nil && k > 0 {
		fast = fast.Next
		k--
		ln++
	}
	counter := ln
	
	if k > 0 {
		counter -= k % ln
		for counter > 0 {
			slow = slow.Next
			counter--
		}
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	if head == slow.Next {
		return	head
	}

	fast.Next = head 
	head = slow.Next 
	slow.Next = nil 
	return head

}

// @lc code=end
type ListNode struct {
	Val  int
	Next *ListNode
}
