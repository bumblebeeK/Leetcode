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
	pre := newHead

	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	s := pre.Next //记录内串的左端
	var cur, tmp, tmp2 *ListNode
	for i := 0; i <= right-left; i++ {
		if i == 0 {
			tmp2 = s
		}
		tmp = s.Next
		s.Next = cur
		cur = s
		s = tmp
		if i == right-left {
			tmp2.Next = s  //接上内串的右串
			pre.Next = cur //接上内串的左串
		}
	}

	return newHead.Next
}

// @lc code=end
