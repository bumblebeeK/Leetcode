package solution

import "fmt"

/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil && fast.Next.Next != nil && fast != slow {
		fast, slow = fast.Next.Next, slow.Next
	}
	if fast != slow {
		return nil
	}
	slow = &ListNode{-1, head}
	fmt.Println(slow.Val, fast.Val)
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

// @lc code=end

type ListNode struct {
	Val  int
	Next *ListNode
}

// 相遇走了 a + kL + b = a + (k-1)L - c
