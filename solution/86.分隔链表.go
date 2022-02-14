package solution

/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	newHead := &ListNode{0, head}
	mark := newHead
	for mark != nil {
		if mark.Next != nil && mark.Next.Val >= x {
			break
		}
		mark = mark.Next
	}
	if mark == nil {
		return head
	}
	tmp, last := mark.Next, mark.Next 
	for tmp != nil {
		if tmp.Val < x {
			next := tmp.Next    
			target := mark.Next 
			mark.Next = tmp
			tmp.Next = target
			mark = mark.Next
			tmp = next
			last.Next = tmp
		} else {
			last = tmp
			tmp = tmp.Next
		}
	}
	return newHead.Next

}

// @lc code=end
type ListNode struct {
	Val  int
	Next *ListNode
}
