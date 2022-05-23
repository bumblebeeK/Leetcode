package solution

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	len := len(lists)
	if len < 1 {
		return nil
	}
	if len == 1 {
		return lists[0]
	}
	l := mergeKLists(lists[:len>>1])
	r := mergeKLists(lists[len>>1:])

	return mergeTwoLists(l, r)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}

// @lc code=end
type ListNode struct {
	Val  int
	Next *ListNode
}
