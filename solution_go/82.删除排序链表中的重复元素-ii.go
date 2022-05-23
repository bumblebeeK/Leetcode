package solution

/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
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
	dummyHead := &ListNode{0, head}
	cursor := dummyHead
	for cursor.Next != nil {
		val := cursor.Next.Val
		delFlag := false
		for cursor.Next.Next != nil && cursor.Next.Next.Val == val {
			cursor.Next = cursor.Next.Next
			delFlag = true
		}
		if delFlag {
			cursor.Next = cursor.Next.Next
			continue
		}
		cursor = cursor.Next

	}
	return dummyHead.Next

}

// @lc code=end
type ListNode struct {
	Val  int
	Next *ListNode
}
