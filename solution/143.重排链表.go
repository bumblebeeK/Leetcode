package solution

/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	l := 0
	dummyHead := head
	for dummyHead != nil {
		l++
		dummyHead = dummyHead.Next
	}
	if l == 1 {
		return
	}
	dummyHead = &ListNode{0, head}

	l = l / 2
	for l != 0 {
		dummyHead = dummyHead.Next
		l--
	}
	rHead := dummyHead.Next
	dummyHead.Next = nil
	dummyHead = nil
	for rHead != nil {
		next := rHead.Next
		rHead.Next = dummyHead
		dummyHead = rHead
		rHead = next
	}
	rHead = dummyHead
	dummyHead = head
	for rHead != nil {
		if dummyHead.Next == nil {
			dummyHead.Next = rHead
			return
		}
		rNext := rHead.Next
		lNext := dummyHead.Next
		dummyHead.Next = rHead
		rHead.Next = rHead
		rHead.Next = lNext
		dummyHead = lNext
		rHead = rNext
	}
}

// @lc code=end

type ListNode struct {
	Val  int
	Next *ListNode
}
