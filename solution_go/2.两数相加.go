package solution

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if nil == l1 || nil == l2 {
		return nil
	}

	head := &ListNode{Val: 0, Next: nil}

	cur := head
	carry := 0
	for nil != l1 || nil != l2 {
		var x, y int = 0, 0
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}

		cur.Next = &ListNode{(x + y + carry) % 10, nil}
		cur = cur.Next
		carry = (x + y + carry) / 10

		if nil != l1 {
			l1 = l1.Next
		}
		if nil != l2 {
			l2 = l2.Next
		}
	}
	if carry > 0 {
		cur.Next = &ListNode{carry % 10, nil}
	}
	return head.Next
}

// @lc code=end
