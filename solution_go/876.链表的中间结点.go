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
	l := 0
	cur, target := head, head
	for cur != nil {
		cur = cur.Next
		l++
	}
	l = l / 2
	for l > 0 {
		target = target.Next
		l--
	}

	return target

}

// @lc code=end

type ListNode struct {
	Val  int
	Next *ListNode
}
