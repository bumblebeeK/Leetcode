package solution

/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left >= right {
		return head
	}

	newHead := &ListNode{0, head}
	mark := newHead
	for t := left; t != 1; {
		mark = mark.Next
		t--
	}
	// head, tail := &ListNode{}, &ListNode{}
	var prev, cur, tail, tmp *ListNode
	cur, tail = mark.Next, mark.Next
	for i := left; i <= right; i++ {
		tmp = cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	mark.Next = prev
	tail.Next = cur
	return newHead.Next
}

// @lc code=end
