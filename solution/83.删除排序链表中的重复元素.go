package solution

/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := &ListNode{0, head}
	newHead := cur
	for cur.Next != nil {
		for cur.Next.Next != nil && cur.Next.Next.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return newHead.Next
}

// @lc code=end
type ListNode struct {
	Val  int
	Next *ListNode
}
