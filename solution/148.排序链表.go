package solution

/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middleNode148(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	h1, h2 := head, head
	for h2.Next != nil && h2.Next.Next != nil {
		h1 = h1.Next
		h2 = h2.Next.Next
	}
	return h1
}

func sortList(head *ListNode) *ListNode {
	cur := head
	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}
	if length <= 1 {
		return head
	}
	middle := middleNode148(head)
	cur = middle.Next
	middle.Next = nil
	middle = cur
	l := sortList(head)
	r := sortList(middle)
	return mergeTwoList(l, r)
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoList(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoList(l1, l2.Next)
	return l2

}

// @lc code=end
