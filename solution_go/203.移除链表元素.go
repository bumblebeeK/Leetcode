package solution

/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	newHead := &ListNode{}
	newHead.Next = head
	mark := newHead
	for mark.Next != nil {
		if mark.Next.Val == val {
			mark.Next = mark.Next.Next
		} else {
			mark = mark.Next
		}
	}

	return newHead.Next
}

// @lc code=end
