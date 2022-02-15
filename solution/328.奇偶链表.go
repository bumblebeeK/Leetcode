package solution

/*
 * @lc app=leetcode.cn id=328 lang=golang
 *
 * [328] 奇偶链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	oddDummy := &ListNode{0, nil}  //序号为奇数的
	evenDummy := &ListNode{0, nil} //序号为偶数的
	evenHead, oddHead := evenDummy, oddDummy
	index := 0
	for head != nil {
		index++
		if index%2 == 0 {
			evenDummy.Next = head
			evenDummy = evenDummy.Next
		} else {
			oddDummy.Next = head
			oddDummy = oddDummy.Next
		}
		head = head.Next
	}
	oddDummy.Next = evenHead.Next
	evenDummy.Next = nil
	return oddHead.Next
}

// @lc code=end
type ListNode struct {
	Val  int
	Next *ListNode
}
