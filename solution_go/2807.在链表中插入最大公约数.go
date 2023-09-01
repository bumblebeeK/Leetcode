package solution

/*
 * @lc app=leetcode.cn id=2807 lang=golang
 *
 * [2807] 在链表中插入最大公约数
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	newHead := head

	for head.Next != nil {
		next := head.Next
		head.Next = &ListNode{
			Val:  GCD(head.Val, next.Val),
			Next: next,
		}
		head.Next.Next = next

		head = next
	}

	return newHead
}

func GCD(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	if b > a {
		a, b = b, a
	}
	r := a % b
	for r != 0 {
		a = b
		b = r
		r = a % b
	}

	return b

}

// @lc code=end
type ListNode struct {
	Val  int
	Next *ListNode
}
