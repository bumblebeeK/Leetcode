package solution

/*
 * @lc app=leetcode.cn id=876 lang=golang
 *
 * [876] 链表的中间结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	p1, p2, cur := head, head, head
	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}

	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	if length%2 == 0 {
		return p1.Next
	} else {
		return p1
	}
}

// @lc code=end
