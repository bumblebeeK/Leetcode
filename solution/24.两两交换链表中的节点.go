package solution

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	pre, cursor := &ListNode{0, head}, head
	newHead := pre
	for cursor != nil && cursor.Next != nil {
		tmp := cursor.Next.Next
		next := cursor.Next
		next.Next = cursor
		pre.Next = next
		cursor.Next = tmp
		pre = cursor
		cursor = tmp
	}
	return newHead.Next
}

// @lc code=end
type ListNode struct {
	Val  int
	Next *ListNode
}
