package solution

/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	front := head
	finished := false
	var recursiveCheck func(*ListNode) bool
	recursiveCheck = func(ln *ListNode) bool {
		if finished {
			return true
		}
		if ln != nil {
			if !recursiveCheck(ln.Next) {
				return false
			}

			if front.Next == ln && front.Val == ln.Val || front == ln {
				finished = true
			}

			if ln.Val != front.Val {
				return false
			}
			front = front.Next
		}
		return true
	}
	return recursiveCheck(head)
}

// @lc code=end

type ListNode struct {
	Val  int
	Next *ListNode
}
