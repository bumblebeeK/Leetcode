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
// func partition(head *ListNode, x int) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	newHead := &ListNode{0, head}
// 	mark := newHead
// 	for mark != nil {
// 		if mark.Next != nil && mark.Next.Val >= x {
// 			break
// 		}
// 		mark = mark.Next
// 	}
// 	if mark == nil {
// 		return head
// 	}
// 	after, last := mark.Next, mark.Next
// 	for after != nil {
// 		if after.Val < x {
// 			next := after.Next
// 			target := mark.Next
// 			mark.Next = after
// 			after.Next = target
// 			mark = mark.Next
// 			after = next
// 			last.Next = after
// 		} else {
// 			last = after
// 			after = after.Next
// 		}
// 	}

// 	return newHead.Next

// }

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	ln := partition(head.Next, x)
	if head.Val < x {
		head.Next = ln
		return head
	}
	prev := &ListNode{0, ln}
	newHead := prev
	for ln != nil {
		if ln.Val >= x {
			prev.Next = head
			head.Next = ln
			return newHead.Next
		}
		prev = ln
		ln = ln.Next
	}
	prev.Next = head
	head.Next = nil
	return newHead.Next

}

// @lc code=end
type ListNode struct {
	Val  int
	Next *ListNode
}
